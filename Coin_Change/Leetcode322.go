package main

func coinChange(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	inf := 0x3f3f3f3f
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0
	for i := range coins {
		for j := 0; j <= amount; j++ {
			if j >= coins[i] {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == inf {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
