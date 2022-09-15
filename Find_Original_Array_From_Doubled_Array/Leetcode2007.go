package main

import "sort"

func findOriginalArray(changed []int) []int {

	if len(changed)&1 == 1 {
		return []int{}
	}
	sort.Ints(changed)
	hash := make(map[int]int)
	for _, x := range changed {
		hash[x]++
	}
	res := make([]int, 0)
	for _, x := range changed {
		if hash[x] != 0 {
			if hash[x*2] == 0 {
				return []int{}
			}
			res = append(res, x)
			hash[x]--
			hash[x*2]--
		}
	}
	return res
}
