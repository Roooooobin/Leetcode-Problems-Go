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
	}
	return y
}

/*
- -去重+二分
对排序去重后的nums中的一段区间[l,r], 如果要保留所有元素, 需要替换区间外的元素填充, nums[r]-nums[l]+1-(r-l+1)
区间外元素个数不能少于填充的元素个数
nums[r]-nums[l]+1-(r-l+1) <= n-(r-l+1)
nums[l] >= nums[r]-n+1
枚举nums[r], 得到最小满足该式的l,此时[l,r]能够保留
*/
