package main

import "sort"

func getKth(lo int, hi int, k int) int {

	hash := make([]int, 1007)
	nums := make([]int, 0)
	for i := lo; i <= hi; i++ {
		nums = append(nums, i)
		x := i
		cnt := 0
		for x != 1 {
			if x&1 == 1 {
				x = 3*x + 1
			} else {
				x >>= 1
			}
			cnt++
		}
		hash[i] = cnt
	}
	sort.Slice(nums, func(i, j int) bool {
		ni, nj := nums[i], nums[j]
		if hash[ni] != hash[nj] {
			return hash[ni] < hash[nj]
		} else {
			return ni < nj
		}
	})
	return nums[k-1]
}
