package main

import "math/bits"

func countPrimeSetBits(left, right int) int {

	res := 0
	mask := 665772 // 10100010100010101100
	for x := left; x <= right; x++ {
		if 1<<bits.OnesCount(uint(x))&mask != 0 {
			res++
		}
	}
	return res
}
