package main

import "container/heap"

func lastStoneWeight(stones []int) int {

	pq := hp{}
	for _, stone := range stones {
		heap.Push(&pq, stone)
	}
	for len(pq) >= 2 {
		stone1 := heap.Pop(&pq).(int)
		stone2 := heap.Pop(&pq).(int)
		if stone1 > stone2 {
			heap.Push(&pq, stone1-stone2)
		}
	}
	if len(pq) == 0 {
		return 0
	}
	return heap.Pop(&pq).(int)
}

type hp []int

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i] > h[j] }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *hp) Pop() interface{}   { a := *h; n := len(*h); v := a[n-1]; *h = a[:n-1]; return v }
