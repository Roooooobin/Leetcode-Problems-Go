package main

func longestArithSeqLength(nums []int) int {

	n := len(nums)
	dp := make([]int, n*n)
	res := 0
	for i := 1; i < n; i++ {
		mp := make(map[int]int)
		for j := 0; j < i; j++ {
			if idx, ok := mp[2*nums[j]-nums[i]]; ok {
				dp[j*n+i] = max(dp[j*n+i], dp[idx*n+j]+1)
				res = max(dp[j*n+i]+2, res)
			}
			mp[nums[j]] = j
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
- -DP
与最长fib数列类似, dp[j][i] = max(dp[j][i], dp[idx][j] + 1), idx用mp记录, mp[2*nums[j]-nums[i]] = idx
*/
