package main

func longestPalindrome(s string) string {

	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	res := 1
	start := 0
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if j == i+1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}
			if dp[i][j] && (j-i+1 > res) {
				res = j - i + 1
				start = i
			}
		}
	}
	return s[start : start+res]
}

/*
- -DP
dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
*/
