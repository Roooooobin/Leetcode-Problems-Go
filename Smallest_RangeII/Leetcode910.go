package main

import "sort"

func smallestRangeII(nums []int, k int) int {

	sort.Ints(nums)
	n := len(nums)
	res := nums[n-1] - nums[0]
	for i := 0; i < n-1; i++ {
		x, y := nums[i], nums[i+1]
		hi := max(x+k, nums[n-1]-k)
		lo := min(y-k, nums[0]+k)
		res = min(res, hi-lo)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {

	if x < y {
		return x
	}
	return y
}
