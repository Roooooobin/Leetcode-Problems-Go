package main

func maximumDifference(nums []int) int {

	n := len(nums)
	res := -1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] < nums[j] {
				if res < nums[j]-nums[i] {
					res = nums[j] - nums[i]
				}
			}
		}
	}
	return res
}
