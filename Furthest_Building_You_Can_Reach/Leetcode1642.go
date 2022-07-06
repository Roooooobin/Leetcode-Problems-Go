package main

import "container/heap"

func furthestBuilding(heights []int, bricks int, ladders int) int {

	n := len(heights)
	diff := make([]int, 0)
	for i := 0; i < n-1; i++ {
		if heights[i+1] > heights[i] {
			diff = append(diff, heights[i+1]-heights[i])
		}
	}
	// just use ladders
	if ladders >= len(diff) {
		return n - 1
	}
	h := &hp{}
	heap.Init(h)
	useBrick := 0
	for _, x := range diff {
		if h.Len() < ladders {
			heap.Push(h, x)
			continue
		}
		top := 0x3f3f3f3f
		if h.Len() > 0 {
			top = heap.Pop(h).(int)
		}
		if x < top {
			bricks -= x
			if bricks < 0 {
				heap.Push(h, top)
				break
			}
			heap.Push(h, top)
			useBrick++
		} else {
			bricks -= top
			if bricks < 0 {
				heap.Push(h, top)
				break
			}
			useBrick++
			heap.Push(h, x)
		}
	}
	canSkip := min(h.Len(), ladders) + useBrick
	for i := 0; i < n-1; i++ {
		if heights[i+1] > heights[i] {
			if canSkip > 0 {
				canSkip--
			} else {
				return i
			}
		}
	}
	return n - 1
}

type hp []int

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i] < h[j] }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type maxSlice []int

func (h maxSlice) Len() int           { return len(h) }
func (h maxSlice) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h maxSlice) Less(i, j int) bool { return h[i] > h[j] }
func (h *maxSlice) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxSlice) Pop() interface{} {
	old := *h
	res := old[len(old)-1]
	*h = old[:len(old)-1]
	return res
}

func furthestBuildingBetter(heights []int, bricks int, ladders int) int {
	h := &maxSlice{}
	heap.Init(h)
	for i := 0; i < len(heights)-1; i++ {
		diff := heights[i+1] - heights[i]
		if diff > 0 {
			heap.Push(h, diff)
			bricks -= diff
			if bricks < 0 {
				if ladders > 0 {
					ladders--
					bricks += heap.Pop(h).(int)
				}
				if bricks < 0 {
					return i
				}
			}
		}
	}
	return len(heights) - 1
}

/*
- -贪心
用大顶堆保存要用ladder, 每次都先用bricks尝试, 如果不够,放入堆中, 如果还有ladder,那么用大顶堆中最大的放回bricks, 如果没有ladder了或是用ladder也还不够,则结束
*/
