package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodeHeap []*ListNode

func (h NodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h NodeHeap) Len() int {
	return len(h)
}

func (h *NodeHeap) Push(v interface{}) {
	*h = append(*h, v.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func mergeKLists(lists []*ListNode) *ListNode {

	hp := &NodeHeap{}
	for _, list := range lists {
		if list == nil {
			continue
		}
		heap.Push(hp, list)
	}
	heap.Init(hp)
	res := &ListNode{0, nil}
	p := res
	for hp.Len() > 0 {
		top := heap.Pop(hp).(*ListNode)
		p.Next = &ListNode{top.Val, nil}
		p = p.Next
		if top.Next != nil {
			top = top.Next
			heap.Push(hp, top)
		}
	}
	return res.Next
}
