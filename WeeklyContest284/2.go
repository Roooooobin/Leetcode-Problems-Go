package main

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {

	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	for _, d := range dig {
		grid[d[0]][d[1]] = 1
	}
	res := 0
	for _, artifact := range artifacts {
		flag := true
		for i := artifact[0]; i <= artifact[2]; i++ {
			for j := artifact[1]; j <= artifact[3]; j++ {
				if grid[i][j] != 1 {
					flag = false
					break
				}
			}
		}
		if flag {
			res++
		}
	}
	return res
}
