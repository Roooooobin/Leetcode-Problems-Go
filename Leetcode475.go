package main

import "sort"

func findRadius(houses []int, heaters []int) int {

	sort.Ints(houses)
	sort.Ints(heaters)
	lo := 0
	hi := 0x3f3f3f3f
	for lo <= hi {
		r := lo + ((hi - lo) >> 1)
		if checkRadius(houses, heaters, r) {
			hi = r - 1
		} else {
			lo = r + 1
		}
	}
	return lo
}

func checkRadius(houses []int, heaters []int, r int) bool {

	n := len(houses)
	i := 0
	for _, heater := range heaters {
		lo := heater - r
		hi := heater + r
		for i < n {
			if houses[i] >= lo && houses[i] <= hi {
				i++
			} else {
				break
			}
		}
	}
	return i >= n
}
