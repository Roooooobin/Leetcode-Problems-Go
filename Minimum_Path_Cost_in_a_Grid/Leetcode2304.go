package main

func minPathCost(grid [][]int, moveCost [][]int) int {

	m, n := len(grid), len(grid[0])
	INF := 0x3f3f3f3f
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	for j := 0; j < n; j++ {
		dp[0][j] = grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				dp[i][j] = min(dp[i][j], dp[i-1][k]+moveCost[grid[i-1][k]][j])
			}
			dp[i][j] += grid[i][j]
		}
	}
	res := INF
	for j := 0; j < n; j++ {
		res = min(res, dp[m-1][j])
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
