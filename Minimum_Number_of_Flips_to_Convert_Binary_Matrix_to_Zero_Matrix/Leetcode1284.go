package main

type Node struct {
	a    [][]int
	step int
}

//
//func main() {
//	minFlips([][]int{{0, 0}, {0, 1}})
//}

func minFlips(mat [][]int) int {

	m, n := len(mat), len(mat[0])
	queue := make([]Node, 0)
	seen := make(map[int]bool)
	queue = append(queue, Node{mat, 0})
	seen[calcMask(mat)] = true
	directions := [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		cur, step := front.a, front.step
		if calcMask(cur) == 0 {
			return step
		}
		for i := range cur {
			for j := range cur[i] {
				nxt := make([][]int, m)
				for k := 0; k < m; k++ {
					nxt[k] = make([]int, n)
					for l := 0; l < n; l++ {
						nxt[k][l] = cur[k][l]
					}
				}
				nxt[i][j] ^= 1
				for _, direction := range directions {
					ni, nj := i+direction[0], j+direction[1]
					if ni >= 0 && nj >= 0 && ni < m && nj < n {
						nxt[ni][nj] ^= 1
					}
				}
				mask := calcMask(nxt)
				if !seen[mask] {
					queue = append(queue, Node{nxt, step + 1})
					seen[mask] = true
				}
			}
		}
	}
	return -1
}

func calcMask(a [][]int) int {
	mask := 0
	n := len(a[0])
	for i := range a {
		for j := range a[i] {
			if a[i][j] == 1 {
				mask |= 1 << (i*n + j)
			}
		}
	}
	return mask
}
