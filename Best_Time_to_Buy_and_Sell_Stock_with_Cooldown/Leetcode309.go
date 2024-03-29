package main

func maxProfit(prices []int) int {

	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][3]int, n)
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[n-1][1], dp[n-1][2])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
- -DP
dp[i][0]持有股票的最大收益
dp[i][1]不持有且处于冷冻期
dp[i][2]不持有且不处于冷冻期
*/
