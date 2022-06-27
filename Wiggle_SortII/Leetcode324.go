package main

import "sort"

func wiggleSort(nums []int) {

	n := len(nums)
	a := make([]int, n)
	sort.Ints(nums)
	l, r := (n-1)/2, n-1
	for i := range a {
		if i&1 == 0 {
			a[i] = nums[l]
			l--
		} else {
			a[i] = nums[r]
			r--
		}
	}
	for i := range a {
		nums[i] = a[i]
	}
}
