package main

import (
	"math/rand"
	"sort"
)

type Solution struct {
	rectangles [][]int
	sum        []int
}

func Constructor(rects [][]int) Solution {
	sum := make([]int, len(rects)+1)
	for i, r := range rects {
		a, b, x, y := r[0], r[1], r[2], r[3]
		sum[i+1] = sum[i] + (x-a+1)*(y-b+1)
	}
	return Solution{rects, sum}
}

func (s *Solution) Pick() []int {
	k := rand.Intn(s.sum[len(s.sum)-1])
	// find which rectangle
	rectIndex := sort.SearchInts(s.sum, k+1) - 1
	r := s.rectangles[rectIndex]
	a, b, y := r[0], r[1], r[3]
	// get the coordination
	da := (k - s.sum[rectIndex]) / (y - b + 1)
	db := (k - s.sum[rectIndex]) % (y - b + 1)
	return []int{a + da, b + db}
}
