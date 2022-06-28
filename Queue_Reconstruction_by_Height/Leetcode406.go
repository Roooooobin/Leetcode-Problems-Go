package main

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {

	sort.Slice(people, func(i, j int) bool {
		pi, pj := people[i], people[j]
		if pi[0] == pj[0] {
			return pi[1] > pj[1]
		} else {
			return pi[0] < pj[0]
		}
	})
	n := len(people)
	res := make([][]int, n)
	indexes := make([]int, n)
	for i := range indexes {
		indexes[i] = i
	}
	for _, peo := range people {
		idx := peo[1]
		realIdx := indexes[idx]
		indexes = append(indexes[:idx], indexes[idx+1:]...)
		res[realIdx] = peo
	}
	return res
}

//func main() {
//	reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}})
//}

/*
贪心+排序
从低到高排序,身高相同的按照p[1]倒序
先放矮的,同时改变下标数组indexes,始终按照indexes[p[1]]放置
*/
