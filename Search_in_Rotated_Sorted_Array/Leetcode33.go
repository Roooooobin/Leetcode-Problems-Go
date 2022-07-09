package main

func search(nums []int, target int) int {

	n := len(nums)
	lo, hi := 0, n-1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if nums[mi] == target {
			return mi
		}
		if nums[mi] >= nums[lo] {
			if target <= nums[mi] && target >= nums[lo] {
				hi = mi - 1
			} else {
				lo = mi + 1
			}
		} else {
			if target >= nums[mi] && target <= nums[hi] {
				lo = mi + 1
			} else {
				hi = mi - 1
			}
		}
	}
	return -1
}
