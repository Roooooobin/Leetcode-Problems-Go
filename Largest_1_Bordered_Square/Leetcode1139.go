package main

func largest1BorderedSquare(grid [][]int) int {

	m, n := len(grid), len(grid[0])
	left := make([][]int, m+1)
	up := make([][]int, m+1)
	for i := range left {
		left[i] = make([]int, n+1)
		up[i] = make([]int, n+1)
	}
	maxBorder := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if grid[i-1][j-1] == 1 {
				left[i][j] = left[i][j-1] + 1
				up[i][j] = up[i-1][j] + 1
				curBorder := min(left[i][j], up[i][j])
				for left[i-curBorder+1][j] < curBorder || up[i][j-curBorder+1] < curBorder {
					curBorder--
				}
				maxBorder = max(maxBorder, curBorder)
			}
		}
	}
	return maxBorder * maxBorder
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
- -动态规划
枚举向左left连续的1个数 和 向上up连续的1个数
*/
