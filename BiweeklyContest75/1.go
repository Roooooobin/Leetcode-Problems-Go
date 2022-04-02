package main

func minBitFlips(start int, goal int) int {

	res := 0
	for start != 0 || goal != 0 {
		x := start % 2
		y := goal % 2
		if x != y {
			res++
		}
		start /= 2
		goal /= 2
	}
	return res
}
