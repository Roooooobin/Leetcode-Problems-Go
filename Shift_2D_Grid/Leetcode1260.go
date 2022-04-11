package main

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	k = k % (m * n)

	flatten := make([]int, 0)
	for _, row := range grid {
		for _, v := range row {
			flatten = append(flatten, v)
		}
	}
	flatten = append(flatten[len(flatten)-k:], flatten[:len(flatten)-k]...)
	idx := 0
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = flatten[idx]
			idx++
		}
	}
	return grid
}
