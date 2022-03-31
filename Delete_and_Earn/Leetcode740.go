package Delete_and_Earn

func deleteAndEarn(nums []int) int {

	a := make([]int, 10003)
	for _, num := range nums {
		a[num] += num
	}
	dp := make([]int, 10003)
	dp[1] = a[1]
	for i := 2; i < len(a); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+a[i])
	}
	return dp[len(a)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
