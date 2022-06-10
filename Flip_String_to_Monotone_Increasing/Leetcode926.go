package main

func minFlipsMonoIncr(s string) int {

	dp := make([]int, 2)
	for _, c := range s {
		dp[1] = min(dp[0], dp[1]) + 1 - int(c-'0')
		dp[0] = dp[0] + int(c-'0')
	}
	return min(dp[0], dp[1])
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
