package main

import (
	"container/heap"
	"github.com/emirpasic/gods/trees/redblacktree"
)

/*
https://leetcode-cn.com/problems/find-servers-that-handled-most-number-of-requests/solution/zhao-dao-chu-li-zui-duo-qing-qiu-de-fu-w-e0a5/
*/
func busiestServers(k int, arrival, load []int) (res []int) {

	available := redblacktree.NewWithIntComparator()
	for id := 0; id < k; id++ {
		available.Put(id, nil)
	}
	busy := hp{}
	requests := make([]int, k)
	for i, t := range arrival {
		for len(busy) > 0 && busy[0].end <= t {
			available.Put(busy[0].id, nil)
			heap.Pop(&busy)
		}
		// dump this request
		if available.Size() == 0 {
			continue
		}
		node, _ := available.Ceiling(i % k)
		if node == nil {
			node = available.Left()
		}
		id := node.Key.(int)
		requests[id]++
		heap.Push(&busy, pair{t + load[i], id})
		available.Remove(id)
	}
	maxRequest := 0
	for _, r := range requests {
		if r > maxRequest {
			maxRequest = r
		}
	}
	for i, r := range requests {
		if r == maxRequest {
			res = append(res, i)
		}
	}
	return
}

type pair struct{ end, id int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
