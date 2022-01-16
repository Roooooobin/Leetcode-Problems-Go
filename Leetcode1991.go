package main

func findMiddleIndex2(nums []int) int {

	n := len(nums)
	prefixSum := make([]int, n+1)
	sum := 0
	for i, x := range nums {
		sum += x
		prefixSum[i+1] = prefixSum[i] + x
	}
	for i := 1; i <= n; i++ {
		if prefixSum[i-1] == sum-prefixSum[i] {
			return i - 1
		}
	}
	return -1
}

func findMiddleIndex(nums []int) int {

	sum := 0
	for _, x := range nums {
		sum += x
	}
	leftSum := 0
	for i, x := range nums {
		if leftSum == sum-leftSum-x {
			return i
		}
		leftSum += x
	}
	return -1
}
