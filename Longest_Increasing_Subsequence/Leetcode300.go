package main

import "sort"

func lengthOfLIS(nums []int) (res int) {

	n := len(nums)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res + 1
}

func lengthOfLISBetter(nums []int) int {

	a := make([]int, 0)
	a = append(a, nums[0])
	for _, num := range nums {
		if num > a[len(a)-1] {
			a = append(a, num)
		} else {
			idx := sort.SearchInts(a, num)
			a[idx] = num
		}
	}
	return len(a)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
