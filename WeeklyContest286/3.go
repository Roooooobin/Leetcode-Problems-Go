package main

import "math"

func kthPalindrome(queries []int, intLength int) []int64 {

	n := len(queries)
	res := make([]int64, n)
	base := int(math.Pow10((intLength - 1) / 2))
	maximum := 9 * base
	for i, q := range queries {
		if q > maximum {
			res[i] = -1
			continue
		}
		// left half
		v := base + q - 1
		x := v
		if intLength&1 == 1 {
			x /= 10
		}
		for ; x > 0; x /= 10 {
			v = v*10 + x%10
		}
		res[i] = int64(v)
	}
	return res
}
