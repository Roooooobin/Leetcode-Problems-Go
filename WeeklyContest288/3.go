package main

import (
	"container/heap"
)

func maximumProduct(nums []int, k int) int {

	h := &hp{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
	}
	for k > 0 {
		k--
		min := heap.Pop(h).(int)
		min++
		heap.Push(h, min)
	}
	MOD := int64(1000000007)
	res := int64(1)
	for h.Len() > 0 {
		x := heap.Pop(h).(int)
		res = res * int64(x) % MOD
	}
	return int(res)
}

type hp []int

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i] < h[j] }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *hp) Pop() interface{}   { a := *h; n := len(*h); v := a[n-1]; *h = a[:n-1]; return v }
