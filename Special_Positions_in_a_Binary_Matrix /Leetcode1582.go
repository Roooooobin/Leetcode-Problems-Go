package main

func numSpecial(mat [][]int) (res int) {

	m, n := len(mat), len(mat[0])
	row := make([]int, m)
	col := make([]int, n)
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 1 {
				row[i]++
				col[j]++
			}
		}
	}
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 1 && row[i] == 1 && col[j] == 1 {
				res++
			}
		}
	}
	return res
}
