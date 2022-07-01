package main

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {

	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	res := 0
	for _, box := range boxTypes {
		res += min(box[0], truckSize) * box[1]
		truckSize -= box[0]
		if truckSize <= 0 {
			break
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
