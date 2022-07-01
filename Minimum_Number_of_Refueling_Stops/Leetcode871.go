package main

import (
	"container/heap"
	"fmt"
)

func minRefuelStops(target int, startFuel int, stations [][]int) int {

	if startFuel >= target {
		return 0
	}
	if startFuel < target && len(stations) == 0 || startFuel < stations[0][0] {
		return -1
	}
	res := 0
	cur := startFuel
	hp := &maxHeap{}
	heap.Init(hp)
	i := 0
	n := len(stations)
	for cur < target {
		if cur >= target {
			return res
		}
		for i < n && stations[i][0] <= cur {
			heap.Push(hp, stations[i][1])
			i++
		}
		if hp.Len() == 0 {
			break
		}
		fuel := heap.Pop(hp).(int)
		res++
		cur += fuel
	}
	if cur < target {
		return -1
	}
	return res
}

type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *maxHeap) Pop() interface{}   { a := *h; n := len(*h); v := a[n-1]; *h = a[:n-1]; return v }

func main() {
	fmt.Println(minRefuelStops(1000000, 70768, [][]int{{12575, 171159},
		{81909, 101253}, {163732, 164401}, {190025, 65493},
		{442889, 31147}, {481202, 166081}, {586028, 206379},
		{591952, 52748}, {595013, 9163}, {611883, 217156}}))
}

/*
- -贪心+heap
只要目前的油能到加油站就先拿上, 不够了再加当前最多的油, 还是不够则return -1
*/
