package main

import "fmt"

func maximumCandies(candies []int, k int64) int {

	sum := int64(0)
	for _, candy := range candies {
		sum += int64(candy)
	}
	if sum < k {
		return 0
	}
	lo, hi := 1, 100000000
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		cnt := int64(0)
		for _, candy := range candies {
			cnt += int64(candy / mid)
		}
		if cnt >= k {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return hi
}

func main() {
	fmt.Println(maximumCandies([]int{5, 8, 7, 7}, 3))
}
