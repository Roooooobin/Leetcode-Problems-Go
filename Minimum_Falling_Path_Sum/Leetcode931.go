package main

func minFallingPathSum(matrix [][]int) int {

	m, n := len(matrix), len(matrix[0])
	dp := make([]int, n)
	for j := 0; j < n; j++ {
		dp[j] = 0x3f3f3f3f
	}
	for j := 0; j < n; j++ {
		dp[j] = matrix[0][j]
	}
	for i := 1; i < m; i++ {
		cur := make([]int, n)
		for j := 0; j < n; j++ {
			cur[j] = dp[j] + matrix[i][j]
			if j > 0 {
				cur[j] = min(cur[j], dp[j-1]+matrix[i][j])
			}
			if j < n-1 {
				cur[j] = min(cur[j], dp[j+1]+matrix[i][j])
			}
		}
		dp = cur
	}
	res := 0x3f3f3f3f
	for j := 0; j < n; j++ {
		res = min(res, dp[j])
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
- -DP
一维滚动数组优化空间复杂度
*/
