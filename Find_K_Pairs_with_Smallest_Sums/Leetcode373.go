package main

import "container/heap"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	n1, n2 := len(nums1), len(nums2)
	hp := &maxHeap{nil, nums1, nums2}
	for i := 0; i < n1 && i < k; i++ {
		hp.data = append(hp.data, Node{i, 0})
	}
	res := make([][]int, 0)
	for hp.Len() > 0 && len(res) < k {
		top := heap.Pop(hp).(Node)
		i, j := top.i, top.j
		res = append(res, []int{nums1[i], nums2[j]})
		if j+1 < n2 {
			heap.Push(hp, Node{i, j + 1})
		}
	}
	return res
}

type Node struct {
	i, j int
}
type maxHeap struct {
	data         []Node
	nums1, nums2 []int
}

func (h maxHeap) Len() int { return len(h.data) }
func (h maxHeap) Less(i, j int) bool {
	return h.nums1[h.data[i].i]+h.nums2[h.data[i].j] < h.nums1[h.data[j].i]+h.nums2[h.data[j].j]
}
func (h maxHeap) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *maxHeap) Push(v interface{}) { h.data = append(h.data, v.(Node)) }
func (h *maxHeap) Pop() interface{} {
	a := h.data
	n := len(a)
	v := a[n-1]
	h.data = a[:n-1]
	return v
}

/*
- -多路归并
用小顶堆找到当前最小的下标,然后拓展,前k个一定会在k次拓展中得出
*/
