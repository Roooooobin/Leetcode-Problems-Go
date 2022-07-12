package main

import (
	"container/heap"
	"strconv"
	"strings"
)

/*
https://leetcode.cn/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/solution/go-xiao-ding-dui-shuang-bai-by-flare-dlcmc/
*/
func kthSmallest(a [][]int, K int) int {

	m, n := len(a), len(a[0])
	h := &Nodes{}
	heap.Init(h)
	curSum := 0
	for i := 0; i < m; i++ {
		curSum += a[i][0]
	}
	heap.Push(h, Node{indexes: make([]int, m), sum: curSum})
	seen := make(map[string]struct{})
	for k := 1; k < K; k++ {
		top := heap.Pop(h).(Node)
		curSum = top.sum
		for i, j := range top.indexes {
			if j < n-1 {
				newIndexes := make([]int, len(top.indexes))
				copy(newIndexes, top.indexes)
				// forward
				newIndexes[i]++
				key := makeKey(newIndexes)
				if _, ok := seen[key]; !ok {
					newSum := curSum + a[i][j+1] - a[i][j]
					heap.Push(h, Node{newIndexes, newSum})
					seen[key] = struct{}{}
				}
			}
		}
	}
	return heap.Pop(h).(Node).sum
}

func makeKey(arr []int) string {
	var builder strings.Builder
	for _, x := range arr {
		builder.WriteString(strconv.Itoa(x))
	}
	return builder.String()
}

type Node struct {
	indexes []int
	sum     int
}

type Nodes []Node

func (h Nodes) Len() int            { return len(h) }
func (h Nodes) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h Nodes) Less(i, j int) bool  { return h[i].sum < h[j].sum }
func (h *Nodes) Push(x interface{}) { *h = append(*h, x.(Node)) }
func (h *Nodes) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

type hp []int

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i] < h[j] }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *hp) Pop() interface{} {
	a := *h
	x := a[len(a)-1]
	*h = a[:len(a)-1]
	return x
}
