package main

import "sort"

func bestSeqAtIndex(height []int, weight []int) int {

	idxes := make([]int, len(height))
	for i := range idxes {
		idxes[i] = i
	}
	sort.Slice(idxes, func(i, j int) bool {
		i, j = idxes[i], idxes[j]
		if height[i] != height[j] {
			return height[i] < height[j]
		} else {
			return weight[i] > weight[j]
		}
	})
	a := make([]int, len(idxes))
	for i, idx := range idxes {
		a[i] = weight[idx]
	}
	return lengthOfLIS(a)
}

func lengthOfLIS(nums []int) int {

	a := make([]int, 0)
	a = append(a, nums[0])
	for _, num := range nums {
		if num > a[len(a)-1] {
			a = append(a, num)
		} else {
			idx := sort.SearchInts(a, num)
			a[idx] = num
		}
	}
	return len(a)
}
