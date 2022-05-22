package main

func countSubstrings(s string) int {

	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if j == i+1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if dp[i][j] {
				res++
			}
		}
	}
	return res
}
