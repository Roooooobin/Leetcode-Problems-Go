package main

import "sort"

func diagonalSort(a [][]int) [][]int {

	m, n := len(a), len(a[0])
	diag := make(map[int][]int, m+n-1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diag[i-j] = append(diag[i-j], a[i][j])
		}
	}
	for _, v := range diag {
		sort.Ints(v)
	}
	it := make(map[int]int, m+n-1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			a[i][j] = diag[i-j][it[i-j]]
			it[i-j]++
		}
	}
	return a
}

/*
- -排序
对角线下标: i-j
*/
