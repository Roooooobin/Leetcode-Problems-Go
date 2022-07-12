package main

func lastStoneWeightII(stones []int) int {

	n := len(stones)
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	target := sum / 2
	dp := make([]int, target+1)
	for i := 0; i < n; i++ {
		x := stones[i]
		for j := target; j >= x; j-- {
			dp[j] = max(dp[j], dp[j-x]+x)
		}
	}
	return abs(sum - dp[target] - dp[target])
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

/*
- -DP
转化为target = sum / 2的背包问题
*/
