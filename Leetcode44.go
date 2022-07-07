package main

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				// *要么重复dp[i][j-1]，要么为空dp[i-1][j]
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
				// 单个字符本来就相等或者p中的是'?'
			} else if p[j-1] == '?' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

/*
- -DP
dp[i-1][j],表示*代表是空字符,例如ab,ab*
dp[i][j-1],表示*代表非空任何字符,例如abcd,ab* # 附上自顶向下方法
*/
