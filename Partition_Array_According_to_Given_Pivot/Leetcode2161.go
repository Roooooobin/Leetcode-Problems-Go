package main

func pivotArray(nums []int, pivot int) []int {

	a1 := make([]int, 0)
	a2 := make([]int, 0)
	a3 := make([]int, 0)
	for _, num := range nums {
		if num < pivot {
			a1 = append(a1, num)
		} else if num == pivot {
			a2 = append(a2, num)
		} else {
			a3 = append(a3, num)
		}
	}
	for _, num := range a2 {
		a1 = append(a1, num)
	}
	for _, num := range a3 {
		a1 = append(a1, num)
	}
	return a1
}
