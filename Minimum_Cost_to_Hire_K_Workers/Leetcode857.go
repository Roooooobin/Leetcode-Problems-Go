package main

import (
	"container/heap"
	"math"
	"sort"
)

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {

	a := make([]int, len(quality))
	for i := range a {
		a[i] = i
	}
	sort.Slice(a, func(i, j int) bool {
		x, y := a[i], a[j]
		return quality[x]*wage[y] > quality[y]*wage[x]
	})
	total := 0
	h := hp{}
	for i := 0; i < k-1; i++ {
		total += quality[a[i]]
		heap.Push(&h, quality[a[i]])
	}
	res := 1e9
	for i := k - 1; i < len(quality); i++ {
		idx := a[i]
		total += quality[idx]
		heap.Push(&h, quality[idx])
		res = math.Min(res, float64(wage[idx])/float64(quality[idx])*float64(total))
		total -= heap.Pop(&h).(int)
	}
	return res
}

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

/*
- -贪心+排序+优先队列
找权重最高的工资组的最小工资开销
https://leetcode.cn/problems/minimum-cost-to-hire-k-workers/solution/gu-yong-k-ming-gong-ren-de-zui-di-cheng-rsz3t/
*/
