package main

func subArrayRanges(nums []int) int64 {

	res := int64(0)
	n := len(nums)
	for i := 0; i < n; i++ {
		minV := nums[i]
		maxV := nums[i]
		for j := i + 1; j < n; j++ {
			minV = min(minV, nums[j])
			maxV = max(maxV, nums[j])
			res += int64(maxV - minV)
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
