package main

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {

	INF := int(1e9)
	dp := make([][]int, k+2)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	dp[0][src] = 0
	for i := 1; i <= k+1; i++ {
		dp[i][src] = 0
		for _, flight := range flights {
			dp[i][flight[1]] = min(dp[i][flight[1]], dp[i-1][flight[0]]+flight[2])
		}
	}
	if dp[k+1][dst] >= INF {
		return -1
	} else {
		return dp[k+1][dst]
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
- -图
Bellford算法
*/
