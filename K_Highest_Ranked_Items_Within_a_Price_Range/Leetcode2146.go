package main

import (
	"container/heap"
	"fmt"
)

func highestRankedKItems(a [][]int, pricing []int, start []int, k int) [][]int {

	m, n := len(a), len(a[0])
	queue := make([][]int, 0)
	seen := make(map[int]struct{})
	x, y := start[0], start[1]
	queue = append(queue, []int{x*n + y, 0})
	seen[x*n+y] = struct{}{}
	h := &hp{}
	heap.Init(h)
	if a[x][y] >= pricing[0] && a[x][y] <= pricing[1] {
		heap.Push(h, Node{0, a[x][y], x, y})
	}
	directions := [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	for len(queue) > 0 {
		front := queue[0]
		x, y = front[0]/n, front[0]%n
		queue = queue[1:]
		for _, direction := range directions {
			nx, ny := x+direction[0], y+direction[1]
			_, ok := seen[nx*n+ny]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && a[nx][ny] != 0 && !ok {
				seen[nx*n+ny] = struct{}{}
				queue = append(queue, []int{nx*n + ny, front[1] + 1})
				if a[nx][ny] >= pricing[0] && a[nx][ny] <= pricing[1] {
					if h.Len() >= k {
						top := heap.Pop(h).(Node)
						heap.Push(h, top)
						if front[1]+1 > top.distance {
							// stop searching, pop
							queue = queue[:len(queue)-1]
						} else {
							heap.Push(h, Node{
								distance: front[1] + 1,
								price:    a[nx][ny],
								rowNo:    nx,
								colNo:    ny,
							})
							heap.Pop(h)
						}
					} else {
						heap.Push(h, Node{
							distance: front[1] + 1,
							price:    a[nx][ny],
							rowNo:    nx,
							colNo:    ny,
						})
					}
				}
			}
		}
	}
	length := h.Len()
	if length > k {
		length = k
	}
	res := make([][]int, length)
	for i := 0; i < length; i++ {
		top := heap.Pop(h).(Node)
		res[length-i-1] = []int{top.rowNo, top.colNo}
	}
	return res
}

type Node struct {
	distance int
	price    int
	rowNo    int
	colNo    int
}

type hp []Node

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	hi, hj := h[i], h[j]
	if hi.distance != hj.distance {
		return hi.distance > hj.distance
	} else if hi.price != hj.price {
		return hi.price > hj.price
	} else if hi.rowNo != hj.rowNo {
		return hi.rowNo > hj.rowNo
	} else {
		return hi.colNo > hj.colNo
	}
}
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(Node)) }
func (h *hp) Pop() interface{}   { a := *h; n := len(*h); v := a[n-1]; *h = a[:n-1]; return v }

func main() {
	fmt.Println(highestRankedKItems([][]int{{57, 54, 48}, {652, 572, 990}, {632, 199, 306}, {38, 702, 263}, {431, 0, 507},
		{673, 570, 750}, {316, 141, 639}},
		[]int{475, 745}, []int{3, 2}, 4))
	/*
		[[57,54,48],[652,572,990],[632,199,306],[38,702,263],[431,0,507],[673,570,750],[316,141,639]]
		[475,745]
		[3,2]
		4
	*/
}
