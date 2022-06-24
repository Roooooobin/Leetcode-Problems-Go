package main

func kthSmallest(matrix [][]int, k int) int {
	var n, m = len(matrix), len(matrix[0])
	var l, r = matrix[0][0], matrix[n-1][m-1]
	var mid int
	for l < r {
		mid = (l + r) >> 1
		var cnt int
		var j = n - 1
		for i := 0; i < n; i++ {
			for j >= 0 && matrix[i][j] > mid {
				j--
			}
			cnt += j + 1
		}
		if cnt >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
