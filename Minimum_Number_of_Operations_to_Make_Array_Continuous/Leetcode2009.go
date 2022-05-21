package main

import (
	"sort"
)

/*
https://leetcode.cn/problems/minimum-number-of-operations-to-make-array-continuous/solution/on-zuo-fa-by-endlesscheng-l7yi/
*/
func minOperations(nums []int) int {

	n := len(nums)
	sort.Ints(nums)
	nums = getUnique(nums)
	res := 0
	for r, num := range nums {
		l := sort.SearchInts(nums[:r], num-n+1)
		res = max(res, r-l+1)
	}
	return n - res
}

func getUnique(nums []int) []int {

	a := make([]int, 0)
	a = append(a, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			a = append(a, nums[i])
		}
	}
	return a
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	minOperations([]int{4, 2, 5, 3})
}
