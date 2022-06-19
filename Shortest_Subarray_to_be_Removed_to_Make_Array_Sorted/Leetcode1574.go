package main

import "sort"

func findLengthOfShortestSubarray(arr []int) int {

	n := len(arr)
	res := 1
	l := 0
	for l < n-1 {
		if arr[l+1] < arr[l] {
			break
		}
		l++
	}
	if l == n-1 {
		return 0
	}
	res = max(res, l)
	r := n - 1
	for r > 0 {
		if arr[r-1] > arr[r] {
			break
		}
		r--
	}
	if r == 0 {
		return 0
	}
	res = max(res, n-r)
	a := arr[r:]
	for i := 0; i <= l; i++ {
		cur := i + 1 + len(a) - sort.SearchInts(a, arr[i])
		res = max(res, cur)
	}
	return n - res
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	//findLengthOfShortestSubarray([]int{1, 4, 3, 2, 1, 2, 3})
	findLengthOfShortestSubarray([]int{0, 1})
}
