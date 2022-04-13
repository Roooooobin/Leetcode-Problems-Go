package main

func canPartition(nums []int) bool {

	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	inf := 0x3f3f3f3f
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0
	for _, num := range nums {
		for i := target; i >= 0; i-- {
			if i >= num {
				dp[i] = min(dp[i], dp[i-num]+1)
			}
		}
	}
	if dp[target] == inf {
		return false
	}
	return true
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
