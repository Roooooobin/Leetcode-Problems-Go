package main

import "container/heap"

func isPossible(target []int) bool {

	sum := 0
	hp := &maxHeap{}
	for _, x := range target {
		sum += x
		heap.Push(hp, x)
	}
	n := len(target)
	if n == 1 {
		return sum == 1
	}
	heap.Init(hp)
	for sum != n {
		top := heap.Pop(hp).(int)
		otherSum := sum - top
		secondTop := heap.Pop(hp).(int)
		y := (top - secondTop) / otherSum
		if (top-secondTop)%otherSum != 0 || top == secondTop {
			y++
		}
		newX := top - y*otherSum
		if newX <= 0 {
			return false
		}
		sum = newX + otherSum
		heap.Push(hp, secondTop)
		heap.Push(hp, newX)
	}
	return true
}

type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *maxHeap) Pop() interface{}   { a := *h; n := len(*h); v := a[n-1]; *h = a[:n-1]; return v }
