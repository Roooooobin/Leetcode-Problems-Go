package main

import "sort"

func largestSubmatrix(a [][]int) (res int) {

	m, n := len(a), len(a[0])
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] == 1 {
				a[i][j] += a[i-1][j]
			}
		}
	}
	for i := 0; i < m; i++ {
		sort.Ints(a[i])
		for j := n - 1; j >= 0; j-- {
			res = max(res, a[i][j]*(n-j))
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
- -贪心+排序
先预处理得到每列以该点为结尾上面连续的1个数, 然后每行处理时先排序, 更新最大长度*高度
*/
