package main

func fallingSquares(positions [][]int) []int {

	n := len(positions)
	h := make([]int, n)
	for i, p := range positions {
		l1, r1 := p[0], p[0]+p[1]-1
		h[i] = p[1]
		for j := 0; j < i; j++ {
			l2, r2 := positions[j][0], positions[j][0]+positions[j][1]-1
			if l1 <= r2 && r1 >= l2 {
				h[i] = max(h[i], h[j]+p[1])
			}
		}
	}
	for i := 1; i < n; i++ {
		h[i] = max(h[i], h[i-1])
	}
	return h
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
