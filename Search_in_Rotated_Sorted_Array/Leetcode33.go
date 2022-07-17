package main

func search(nums []int, target int) int {

	n := len(nums)
	lo, hi := 0, n-1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if nums[mi] == target {
			return mi
		}
		if nums[mi] >= nums[0] {
			if target < nums[mi] && target >= nums[0] {
				hi = mi - 1
			} else {
				lo = mi + 1
			}
		} else {
			if target > nums[mi] && target <= nums[n-1] {
				lo = mi + 1
			} else {
				hi = mi - 1
			}
		}
	}
	return -1
}
