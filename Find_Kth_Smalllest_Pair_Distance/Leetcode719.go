package main

import "sort"

func smallestDistancePair(nums []int, K int) int {

	n := len(nums)
	sort.Ints(nums)
	lo, hi := 0, nums[n-1]-nums[0]+1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		k := 0
		for i := 0; i < n; i++ {
			target := nums[i] + mi
			k += sort.SearchInts(nums, target) - i - 1
		}
		if k >= K {
			hi = mi - 1
		} else {
			lo = mi + 1
		}
	}
	return lo
}

func main() {
	smallestDistancePair([]int{1, 6, 1}, 3)
}
