package main

func pacificAtlantic(heights [][]int) [][]int {

	m := len(heights)
	n := len(heights[0])
	pac := make([]bool, m*n)
	atl := make([]bool, m*n)
	directions := [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	inArea := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	var dfs func(i, j int, seen []bool)
	dfs = func(i, j int, mark []bool) {
		mark[i*n+j] = true
		for _, direction := range directions {
			ni := i + direction[0]
			nj := j + direction[1]
			if inArea(ni, nj) && !mark[ni*n+nj] {
				if heights[ni][nj] >= heights[i][j] {
					dfs(ni, nj, mark)
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		dfs(i, 0, pac)
		dfs(i, n-1, atl)
	}
	for j := 0; j < n; j++ {
		dfs(0, j, pac)
		dfs(m-1, j, atl)
	}
	res := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pac[i*n+j] && atl[i*n+j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}
