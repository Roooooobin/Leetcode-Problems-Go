package main

import "sort"

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {

	MOD := 1000000007
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	maxH, maxW := horizontalCuts[0], verticalCuts[0]
	horizontalCuts = append(horizontalCuts, h)
	verticalCuts = append(verticalCuts, w)
	for i := 1; i < len(horizontalCuts); i++ {
		maxH = max(maxH, horizontalCuts[i]-horizontalCuts[i-1])
	}
	for i := 1; i < len(verticalCuts); i++ {
		maxW = max(maxW, verticalCuts[i]-verticalCuts[i-1])
	}
	return maxH * maxW % MOD
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
