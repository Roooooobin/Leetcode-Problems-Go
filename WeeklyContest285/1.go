package main

func countHillValley(nums []int) int {

	n := len(nums)
	res := 0
	for i := 1; i < n-1; i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		j := i - 1
		l, r := nums[i], nums[i]
		for ; j >= 0; j-- {
			if nums[j] != nums[i] {
				l = nums[j]
				break
			}
		}
		j = i + 1
		for ; j < n; j++ {
			if nums[j] != nums[i] {
				r = nums[j]
				break
			}
		}
		if l < nums[i] && r < nums[i] {
			res++
		} else if l > nums[i] && r > nums[i] {
			res++
		}
	}
	return res
}
