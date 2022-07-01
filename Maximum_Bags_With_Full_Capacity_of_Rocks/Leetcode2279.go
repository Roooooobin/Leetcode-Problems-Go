package main

import "sort"

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {

	res := 0
	diff := make([]int, 0)
	n := len(capacity)
	for i := 0; i < n; i++ {
		if capacity[i] == rocks[i] {
			res++
		} else {
			diff = append(diff, capacity[i]-rocks[i])
		}
	}
	sort.Ints(diff)
	for _, d := range diff {
		if additionalRocks >= d {
			res++
			additionalRocks -= d
		} else {
			break
		}
	}
	return res
}
