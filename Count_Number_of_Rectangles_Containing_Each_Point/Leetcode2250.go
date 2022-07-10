package main

import "sort"

func countRectangles(rectangles [][]int, points [][]int) []int {

	ys := [101][]int{}
	for i := range ys {
		ys[i] = make([]int, 0)
	}
	for _, rect := range rectangles {
		ys[rect[1]] = append(ys[rect[1]], rect[0])
	}
	for i := range ys {
		sort.Ints(ys[i])
	}
	//binarySearch := func(a []int, target int) int {
	//	n := len(a)
	//	if target < a[0] {
	//		return 0
	//	}
	//	if target > a[n-1] {
	//		return n
	//	}
	//	lo, hi := 0, n-1
	//	for lo <= hi {
	//		mi := lo + (hi-lo)>>1
	//		if target <= a[mi] {
	//			hi = mi - 1
	//		} else {
	//			lo = mi + 1
	//		}
	//	}
	//	return lo
	//}
	res := make([]int, 0)
	for _, point := range points {
		cnt := 0
		for i := point[1]; i <= 100; i++ {
			if len(ys[i]) == 0 {
				continue
			}
			cnt += len(ys[i]) - sort.SearchInts(ys[i], point[0])
		}
		res = append(res, cnt)
	}
	return res
}

/*
- -二分
利用y值范围小, 按y值存储对应的x(按序)+二分
*/
