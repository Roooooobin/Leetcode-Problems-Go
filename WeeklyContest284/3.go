package main

import "sort"

func maximumTop(nums []int, k int) int {

	if k == 0 {
		return nums[0]
	}
	n := len(nums)
	if n == 1 && k%2 == 1 {
		return -1
	} else if n == 1 {
		return nums[0]
	}
	a := make([]int, min(n, k+1))
	for i := 0; i < min(n, k+1); i++ {
		a = append(a, nums[i])
	}
	sort.Ints(a)
	maxV := a[len(a)-1]
	idx := -1
	for i := 0; i < min(n, k+1); i++ {
		if nums[i] == maxV {
			idx = i
			break
		}
	}
	if idx != k-1 {
		return maxV
	} else {
		return a[len(a)-2]
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
