package main

func minimumDeletions(nums []int) int {

	if len(nums) <= 2 {
		return len(nums)
	}
	minm, maxm := 0x3f3f3f3f, -0x3f3f3f3f
	minIdx, maxIdx := -1, -1
	for i, num := range nums {
		if num < minm {
			minm = num
			minIdx = i
		}
		if num > maxm {
			maxm = num
			maxIdx = i
		}
	}
	if minIdx > maxIdx {
		maxIdx, minIdx = minIdx, maxIdx
	}
	return min(min(maxIdx+1, len(nums)-minIdx), len(nums)-(maxIdx-minIdx)+1)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
