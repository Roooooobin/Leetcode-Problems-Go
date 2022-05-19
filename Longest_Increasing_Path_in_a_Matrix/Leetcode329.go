package main

func longestIncreasingPath(a [][]int) int {

	m, n := len(a), len(a[0])
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}
	var dfs func(i, j, val int) int
	dfs = func(i, j, val int) int {
		if i < 0 || i >= m || j < 0 || j >= n {
			return 0
		}
		if val >= a[i][j] {
			return 0
		}
		// visited, return
		if memo[i][j] != 0 {
			return memo[i][j]
		}
		left := dfs(i, j-1, a[i][j])
		right := dfs(i, j+1, a[i][j])
		up := dfs(i+1, j, a[i][j])
		down := dfs(i-1, j, a[i][j])
		memo[i][j] = max(max(max(left, right), up), down) + 1
		return memo[i][j]
	}
	res := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if memo[i][j] == 0 {
				left := dfs(i, j-1, a[i][j])
				right := dfs(i, j+1, a[i][j])
				up := dfs(i+1, j, a[i][j])
				down := dfs(i-1, j, a[i][j])
				memo[i][j] = max(max(max(left, right), up), down) + 1
			}
			res = max(res, memo[i][j])
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
