package main

func minDistance(word1 string, word2 string) int {

	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
			}
		}
	}
	return dp[m][n]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

/*
- -DP
dp[i][j]表示word1[:i]到word2[:j]的距离, 如果word1[i-1] == word2[j-1]那么dp[i][j] = dp[i-1][j-1]
否则需要删除或插入或替换一个字符, dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
*/
