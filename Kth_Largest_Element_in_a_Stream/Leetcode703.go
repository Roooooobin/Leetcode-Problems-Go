package main

import "container/heap"

type KthLargest struct {
	k  int
	pq hp
}

func Constructor(k int, nums []int) KthLargest {
	pq := hp{}
	for _, num := range nums {
		if len(pq) < k {
			heap.Push(&pq, num)
		} else {
			top := heap.Pop(&pq).(int)
			if num > top {
				top = num
			}
			heap.Push(&pq, top)
		}
	}
	return KthLargest{k, pq}
}

func (this *KthLargest) Add(val int) int {

	heap.Push(&this.pq, val)
	if len(this.pq) > this.k {
		heap.Pop(&this.pq)
	}
	top := heap.Pop(&this.pq).(int)
	heap.Push(&this.pq, top)
	return top
}

type hp []int

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i] < h[j] }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *hp) Pop() interface{}   { a := *h; n := len(*h); v := a[n-1]; *h = a[:n-1]; return v }
