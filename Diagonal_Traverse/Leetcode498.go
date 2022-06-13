package main

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	res := make([]int, 0, m*n)
	for i := 0; i < m+n-1; i++ {
		if i&1 == 1 {
			x := max(i-n+1, 0)
			y := min(i, n-1)
			for x < m && y >= 0 {
				res = append(res, mat[x][y])
				x++
				y--
			}
		} else {
			x := min(i, m-1)
			y := max(i-m+1, 0)
			for x >= 0 && y < n {
				res = append(res, mat[x][y])
				x--
				y++
			}
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
