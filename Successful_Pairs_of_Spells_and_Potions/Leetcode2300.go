package main

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {

	n := len(potions)
	res := make([]int, 0)
	sort.Ints(potions)
	binarySearch := func(a []int, t int) int {
		lo, hi := 0, len(a)-1
		t64 := int64(t)
		if t64*int64(a[hi]) < success {
			return len(a)
		}
		if t64*int64(a[lo]) >= success {
			return 0
		}
		for lo <= hi {
			mi := lo + (hi-lo)>>1
			if t64*int64(a[mi]) < success {
				lo = mi + 1
			} else {
				hi = mi - 1
			}
		}
		return lo
	}
	for _, x := range spells {
		res = append(res, n-binarySearch(potions, x))
	}
	return res
}
