package main

func cherryPickup(grid [][]int) int {

	NEGINF := -0x3f3f3f3f
	n := len(grid)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = NEGINF
		}
	}
	dp[0][0] = grid[0][0]
	for k := 1; k < n*2-1; k++ {
		for x1 := min(k, n-1); x1 >= max(k-n+1, 0); x1-- {
			for x2 := min(k, n-1); x2 >= x1; x2-- {
				y1, y2 := k-x1, k-x2
				if grid[x1][y1] == -1 || grid[x2][y2] == -1 {
					dp[x1][x2] = NEGINF
					continue
				}
				res := dp[x1][x2]
				if x1 > 0 {
					res = max(res, dp[x1-1][x2])
				}
				if x2 > 0 {
					res = max(res, dp[x1][x2-1])
				}
				if x1 > 0 && x2 > 0 {
					res = max(res, dp[x1-1][x2-1])
				}
				res += grid[x1][y1]
				if x2 != x1 {
					res += grid[x2][y2]
				}
				dp[x1][x2] = res
			}
		}
	}
	return max(dp[n-1][n-1], 0)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
作者：LeetCode-Solution
链接：https://leetcode.cn/problems/cherry-pickup/solution/zhai-ying-tao-by-leetcode-solution-1h3k/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
- -DP so hard
dp[k][x1][x2]表示两个人分别从(x1, k-x1), (x2, k-x2)同时出发到达(n-1, n-1)摘到樱桃个数之和的最大值
*/
