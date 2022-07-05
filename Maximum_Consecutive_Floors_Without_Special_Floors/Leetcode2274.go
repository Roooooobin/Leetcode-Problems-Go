package main

import "sort"

func maxConsecutive(bottom int, top int, special []int) int {

	sort.Ints(special)
	special = append(special, top+1)
	pre := bottom - 1
	res := 0
	for _, x := range special {
		res = max(res, x-pre-1)
		pre = x
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
