package main

import "sort"

func intersectionSizeTwo(intervals [][]int) (res int) {

	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	n, m := len(intervals), 2
	vals := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		for j, k := intervals[i][0], len(vals[i]); k < m; k++ {
			res++
			for p := i - 1; p >= 0 && intervals[p][1] >= j; p-- {
				vals[p] = append(vals[p], j)
			}
			j++
		}
	}
	return
}

//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/set-intersection-size-at-least-two/solution/she-zhi-jiao-ji-da-xiao-zhi-shao-wei-2-b-vuiv/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
/*
- -贪心
*/
