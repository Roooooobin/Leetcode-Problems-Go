package main

import "sort"

func minSetSize(arr []int) int {

	mp := make(map[int]int)
	for _, num := range arr {
		v := mp[num]
		mp[num] = v + 1
	}
	sort.Slice(arr, func(i, j int) bool {
		x, y := arr[i], arr[j]
		if mp[x] != mp[y] {
			return mp[x] > mp[y]
		} else {
			return x < y
		}
	})
	set := make(map[int]struct{})
	for i := 0; i < (len(arr)+1)/2; i++ {
		set[arr[i]] = struct{}{}
	}
	return len(set)
}
