package main

import (
	"sort"
)

type frac struct {
	x int
	y int
}

func kthSmallestPrimeFraction(arr []int, k int) []int {

	n := len(arr)
	fracs := make([]frac, 0)
	for i, num := range arr {
		for j := i + 1; j < n; j++ {
			fracs = append(fracs, frac{num, arr[j]})
		}
	}
	sort.Slice(fracs, func(i, j int) bool {
		return fracs[i].x*fracs[j].y < fracs[j].x*fracs[i].y
	})
	return []int{fracs[k-1].x, fracs[k-1].y}
}

//func kthSmallestPrimeFraction(arr []int, k int) []int {
//	n := len(arr)
//	h := make(hp, n-1)
//	for j := 1; j < n; j++ {
//		h[j-1] = frac{arr[0], arr[j], 0, j}
//	}
//	heap.Init(&h)
//	for loop := k - 1; loop > 0; loop-- {
//		f := heap.Pop(&h).(frac)
//		if f.i+1 < f.j {
//			heap.Push(&h, frac{arr[f.i+1], f.y, f.i + 1, f.j})
//		}
//	}
//	return []int{h[0].x, h[0].y}
//}

//type frac struct{ x, y, i, j int }
//type hp []frac
//func (h hp) Len() int            { return len(h) }
//func (h hp) Less(i, j int) bool  { return h[i].x*h[j].y < h[i].y*h[j].x }
//func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
//func (h *hp) Push(v interface{}) { *h = append(*h, v.(frac)) }
//func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/k-th-smallest-prime-fraction/solution/di-k-ge-zui-xiao-de-su-shu-fen-shu-by-le-argw/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
