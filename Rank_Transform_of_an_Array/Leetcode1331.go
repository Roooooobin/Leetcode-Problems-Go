package main

import "sort"

func arrayRankTransform(arr []int) []int {

	tmp := make([]int, len(arr))
	copy(tmp, arr)
	sort.Ints(tmp)
	mp := make(map[int]int)
	rank := 0
	for i := range tmp {
		if _, ok := mp[tmp[i]]; ok {
			continue
		} else {
			rank++
			mp[tmp[i]] = rank
		}
	}
	res := make([]int, 0)
	for _, num := range arr {
		res = append(res, mp[num])
	}
	return res
}
