package main

func nthUglyNumber(n int) int {
	nums := []int{1}
	idx2, idx3, idx5 := 0, 0, 0
	for i := 1; i < n; i++ {
		newUgly := min(nums[idx2]*2, min(nums[idx3]*3, nums[idx5]*5))
		nums = append(nums, newUgly)
		if newUgly == nums[idx2]*2 {
			idx2++
		}
		if newUgly == nums[idx3]*3 {
			idx3++
		}
		if newUgly == nums[idx5]*5 {
			idx5++
		}
	}
	return nums[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
