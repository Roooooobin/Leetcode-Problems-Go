package main

func countPalindromicSubsequences(s string) int {

	MOD := int(1e9 + 7)
	dp := [4][][]int{}
	n := len(s)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}
	for i, c := range s {
		dp[c-'a'][i][i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			for k := 0; k < 4; k++ {
				c := byte('a' + k)
				if s[i] == c && s[j] == c {
					dp[k][i][j] = (dp[0][i+1][j-1] + dp[1][i+1][j-1] + dp[2][i+1][j-1] + dp[3][i+1][j-1] + 2) % MOD
				} else if s[i] == c {
					dp[k][i][j] = dp[k][i][j-1]
				} else if s[j] == c {
					dp[k][i][j] = dp[k][i+1][j]
				} else {
					dp[k][i][j] = dp[k][i+1][j-1]
				}
			}
		}
	}
	res := 0
	for _, d := range dp {
		res = (res + d[0][n-1]) % MOD
	}
	return res
}

/*
- -DP
状态dp(x,i,j)表示在字符串区间s[i:j] 中以字符 x 为开头和结尾的不同「回文序列」总数
*/
