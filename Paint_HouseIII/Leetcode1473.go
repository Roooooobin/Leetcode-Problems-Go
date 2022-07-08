package Paint_HouseIII

func minCost(houses []int, cost [][]int, m, n, target int) int {
	const INF = 0x3f3f3f3f
	// 将颜色调整为从 0 开始编号，没有被涂色标记为 -1
	for i := range houses {
		houses[i]--
	}
	// dp 所有元素初始化为极大值
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, target)
			for k := range dp[i][j] {
				dp[i][j][k] = INF
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if houses[i] != -1 && houses[i] != j {
				continue
			}
			for k := 0; k < target; k++ {
				for j0 := 0; j0 < n; j0++ {
					if j == j0 {
						if i == 0 {
							if k == 0 {
								dp[i][j][k] = 0
							}
						} else {
							dp[i][j][k] = min(dp[i][j][k], dp[i-1][j][k])
						}
					} else if i > 0 && k > 0 {
						dp[i][j][k] = min(dp[i][j][k], dp[i-1][j0][k-1])
					}
				}

				if dp[i][j][k] != INF && houses[i] == -1 {
					dp[i][j][k] += cost[i][j]
				}
			}
		}
	}
	res := INF
	for _, a := range dp[m-1] {
		res = min(res, a[target-1])
	}
	if res == INF {
		return -1
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
https://leetcode.cn/problems/paint-house-iii/solution/fen-shua-fang-zi-iii-by-leetcode-solutio-powb/
- -DP
设 dp[i][j][k] 表示将 [0,i] 的房子都涂上颜色，最末尾的第 i 个房子的颜色为 j，并且它属于第k个街区时，需要的最少花费

*/
