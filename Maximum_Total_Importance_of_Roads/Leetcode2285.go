package main

import "sort"

func maximumImportance(n int, roads [][]int) int64 {

	degrees := make([]int, n)
	for _, road := range roads {
		degrees[road[0]]++
		degrees[road[1]]++
	}
	sort.Ints(degrees)
	res := int64(0)
	for i, d := range degrees {
		res += int64(d) * int64(i+1)
	}
	return res
}
