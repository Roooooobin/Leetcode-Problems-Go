package main

import (
	"container/heap"
	"sort"
)

func scheduleCourse(courses [][]int) int {

	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	hp := &maxHeap{}
	heap.Init(hp)
	cur := 0 // start from time: 0
	for _, course := range courses {
		duration, endTime := course[0], course[1]
		if cur+duration <= endTime {
			heap.Push(hp, duration)
			cur += duration
		} else if hp.Len() != 0 {
			top := heap.Pop(hp).(int)
			if top > duration {
				heap.Push(hp, duration)
				cur -= top - duration
			} else {
				heap.Push(hp, top)
			}
		}
	}
	return hp.Len()
}

type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *maxHeap) Pop() interface{}   { a := *h; n := len(*h); v := a[n-1]; *h = a[:n-1]; return v }
