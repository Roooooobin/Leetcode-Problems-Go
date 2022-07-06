package main

import "sort"

func eliminateMaximum(dist []int, speed []int) int {

	for i := range dist {
		dist[i] = (dist[i] - 1) / speed[i]
	}
	sort.Ints(dist)
	for i, d := range dist {
		if i > d {
			return i
		}
	}
	return len(dist)
}

/*
- -贪心+排序
按照到达时间排序, 下标即最晚时间, 如果到达时间小于最晚时间则无法打死该怪
*/
