package main

func kInversePairs(n, k int) int {
	const MOD = 1e9 + 7
	dp := [2][]int{make([]int, k+1), make([]int, k+1)}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			cur := i & 1
			prev := cur ^ 1
			dp[cur][j] = 0
			if j > 0 {
				dp[cur][j] = dp[cur][j-1]
			}
			if j >= i {
				dp[cur][j] -= dp[prev][j-i]
			}
			dp[cur][j] += dp[prev][j]
			if dp[cur][j] >= MOD {
				dp[cur][j] -= MOD
			} else if dp[cur][j] < 0 {
				dp[cur][j] += MOD
			}
		}
	}
	return dp[n&1][k]
}

/*
https://leetcode.cn/problems/k-inverse-pairs-array/solution/gong-shui-san-xie-yi-dao-xu-lie-dp-zhuan-tm01/
- -DP

*/
