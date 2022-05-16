package main

func shortestPathBinaryMatrix(grid [][]int) int {

	if grid[0][0] == 1 {
		return -1
	}
	n := len(grid)
	directions := [8][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}, {-1, -1}, {-1, 1}, {1, 1}, {1, -1}}
	queue := make([][2]int, 0)
	seen := make(map[int]struct{})
	queue = append(queue, [2]int{0, 1})
	seen[0] = struct{}{}
	for len(queue) > 0 {
		front := queue[0][0]
		step := queue[0][1]
		queue = queue[1:]
		x, y := front/n, front%n
		if x == n-1 && y == n-1 {
			return step
		}
		for _, direction := range directions {
			nx, ny := x+direction[0], y+direction[1]
			cord := nx*n + ny
			if nx >= 0 && nx < n && ny >= 0 && ny < n {
				if _, ok := seen[cord]; ok || grid[nx][ny] == 1 {
					continue
				}
				queue = append(queue, [2]int{cord, step + 1})
				seen[cord] = struct{}{}
			}
		}
	}
	return -1
}
