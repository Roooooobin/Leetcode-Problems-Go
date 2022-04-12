package main

func gameOfLife(board [][]int) {

	m, n := len(board), len(board[0])
	numOfLives := make([][]int, m)
	for i := range numOfLives {
		numOfLives[i] = make([]int, n)
	}
	for i := range board {
		for j := range board[i] {
			for k := i - 1; k <= i+1; k++ {
				if k < 0 || k >= m {
					continue
				}
				for l := j - 1; l <= j+1; l++ {
					if l < 0 || l >= n || (k == i && l == j) {
						continue
					}
					if board[k][l] == 1 {
						numOfLives[i][j]++
					}
				}
			}
		}
	}
	for i := range board {
		for j := range board[i] {
			num := numOfLives[i][j]
			if board[i][j] == 1 {
				if num < 2 || num > 3 {
					board[i][j] = 0
				}
			} else {
				if num == 3 {
					board[i][j] = 1
				}
			}
		}
	}
}
