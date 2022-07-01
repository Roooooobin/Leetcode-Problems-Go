package main

func minCut(s string) int {

	n := len(s)
	isPal := make([][]bool, n)
	for i := range isPal {
		isPal[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		isPal[i][i] = true
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if j == i+1 {
				isPal[i][j] = s[i] == s[j]
			} else {
				isPal[i][j] = s[i] == s[j] && isPal[i+1][j-1]
			}
		}
	}
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 0x3f3f3f3f
	}
	for i := 0; i < n; i++ {
		if isPal[0][i] {
			dp[i] = 0
		} else {
			for j := 0; j < i; j++ {
				if isPal[j+1][i] {
					dp[i] = min(dp[i], dp[j]+1)
				}
			}
		}
	}
	return dp[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
- -DP
dp[i]表示0-i之间分割的最小, dp[i] = min(dp[i], dp[j]+1), 如果j+1-i可以分割
*/
