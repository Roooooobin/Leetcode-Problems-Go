package main

/*
https://leetcode.cn/problems/can-i-win/solution/by-ac_oier-0ed9/
*/
func canIWin(n int, t int) bool {

	if t == 0 {
		return true
	}
	if n*(n+1)/2 < t {
		return false
	}
	dp := make([]int, 1<<20)
	var dfs func(state, total int) int
	// 1 true, -1 false
	dfs = func(state, total int) int {
		if dp[state] != 0 {
			return dp[state]
		}
		for i := 0; i < n; i++ {
			// visited
			if (state>>i)&1 == 1 {
				continue
			}
			if total+i+1 >= t {
				dp[state] = 1
				return 1
			}
			if dfs(state|(1<<i), total+i+1) == -1 {
				dp[state] = 1
				return 1
			}
		}
		dp[state] = -1
		return -1
	}
	return dfs(0, 0) == 1
}

/*
- -DP+博弈论+记忆化搜索
dp[state]表示当前选择的掩码mask为state时先手能否赢, 赢的条件时当前选择的total>=t
记忆化dp
*/
