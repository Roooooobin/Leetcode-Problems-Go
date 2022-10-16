package main

//
//import (
//	"container/heap"
//	"fmt"
//	"math/bits"
//)
//
//func main() {
//	var n, k int
//	fmt.Scan(&n)
//	fmt.Scan(&k)
//	a := make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Scan(&a[i])
//	}
//	hp := &maxHeap{}
//	heap.Init(hp)
//	for i := 0; i < n; i++ {
//		heap.Push(hp, Node{a[i], bits.OnesCount(uint(a[i]))})
//	}
//	for i := 0; i < k; i++ {
//		top := heap.Pop(hp).(Node)
//		newVal := top.cnt
//		heap.Push(hp, Node{newVal, bits.OnesCount(uint(newVal))})
//	}
//	res := 0
//	for i := 0; i < n; i++ {
//		top := heap.Pop(hp).(Node)
//		res += top.val
//	}
//	fmt.Println(res)
//}
//
//type Node struct {
//	val int
//	cnt int
//}
//
//type maxHeap []Node
//
//func (h maxHeap) Len() int { return len(h) }
//func (h maxHeap) Less(i, j int) bool {
//	return h[i].val-h[i].cnt > h[j].val-h[j].cnt
//}
//func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
//func (h *maxHeap) Push(v interface{}) { *h = append(*h, v.(Node)) }
//func (h *maxHeap) Pop() interface{}   { a := *h; n := len(*h); v := a[n-1]; *h = a[:n-1]; return v }
