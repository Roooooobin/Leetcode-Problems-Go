package main

import "sort"

func twoCitySchedCost(costs [][]int) int {

	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})
	res := 0
	for i := 0; i < len(costs)/2; i++ {
		res += costs[i][0] + costs[len(costs)-i-1][1]
	}
	return res
}
