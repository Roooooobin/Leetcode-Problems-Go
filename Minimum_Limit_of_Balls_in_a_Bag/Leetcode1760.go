package main

func minimumSize(nums []int, maxOperations int) int {

	lo, hi := 1, 1000000001
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		op := 0
		for _, num := range nums {
			op += (num - 1) / mid
			if op > maxOperations {
				break
			}
		}
		if op > maxOperations {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}

/*
- -二分答案
*/
