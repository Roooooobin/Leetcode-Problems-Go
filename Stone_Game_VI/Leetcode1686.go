package main

import "sort"

func stoneGameVI(aliceValues []int, bobValues []int) int {

	n := len(aliceValues)
	indexes := make([]int, n)
	for i := range indexes {
		indexes[i] = i
	}
	sort.Slice(indexes, func(i, j int) bool {
		i, j = indexes[i], indexes[j]
		return aliceValues[i]+bobValues[i] > aliceValues[j]+bobValues[j]
	})
	scoreAlice := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			scoreAlice += aliceValues[indexes[i]]
		} else {
			scoreAlice -= bobValues[indexes[i]]
		}
	}
	if scoreAlice > 0 {
		return 1
	} else if scoreAlice < 0 {
		return -1
	}
	return 0
}

func main() {
	//stoneGameVI([]int{40, 76, 27, 31, 40, 12, 57, 10, 88, 72, 85, 5, 28, 25, 61, 82, 16, 63, 50, 90, 20, 55, 63}, []int{74, 5, 37, 21, 29, 59, 94, 25, 31, 10, 86, 31, 99, 45, 77, 91, 44, 73, 83, 67, 55, 12, 35})
	stoneGameVI([]int{2, 4, 3}, []int{1, 6, 7})
}
