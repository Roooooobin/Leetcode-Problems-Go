package main

func numSubarrayProductLessThanK(nums []int, k int) int {

	if k == 0 {
		return 0
	}
	res := 0
	n := len(nums)
	curProduct := 1
	l := 0
	for i := 0; i < n; i++ {
		curProduct *= nums[i]
		for curProduct >= k && l <= i {
			curProduct /= nums[l]
			l++
		}
		if curProduct < k {
			res += i - l + 1
		}
	}
	return res
}
