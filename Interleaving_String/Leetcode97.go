package main

func isInterleave(s1 string, s2 string, s3 string) bool {

	m, n := len(s1), len(s2)
	if m == 0 {
		return s2 == s3
	}
	if n == 0 {
		return s1 == s3
	}
	if m+n != len(s3) {
		return false
	}
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			k := i + j - 1
			if i > 0 && s3[k] == s1[i-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j]
			}
			if j > 0 && s3[k] == s2[j-1] {
				dp[i][j] = dp[i][j] || dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}

/*
动态规划
dp[i][j]表示s1[:i]和s2[:j]是否能交织得到s3[:i+j-1]
*/
