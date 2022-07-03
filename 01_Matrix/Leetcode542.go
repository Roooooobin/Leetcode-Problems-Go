package main

func updateMatrix(mat [][]int) [][]int {

	queue := make([]int, 0)
	seen := make(map[int]bool)
	m, n := len(mat), len(mat[0])
	directions := [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 0 {
				cord := i*n + j
				queue = append(queue, cord)
				seen[cord] = true
			}
		}
	}
	step := 1
	for len(queue) > 0 {
		k := len(queue)
		for k > 0 {
			k--
			front := queue[0]
			queue = queue[1:]
			x, y := front/n, front%n
			for _, direction := range directions {
				nx, ny := x+direction[0], y+direction[1]
				ncord := nx*n + ny
				if nx >= 0 && nx < m && ny >= 0 && ny < n && !seen[ncord] {
					seen[ncord] = true
					res[nx][ny] = step
					queue = append(queue, ncord)
				}
			}
		}
		step++
	}
	return res
}
