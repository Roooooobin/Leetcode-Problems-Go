package main

import "sort"

func bagOfTokensScore(a []int, power int) int {

	sort.Ints(a)
	res, points := 0, 0
	i, j := 0, len(a)-1
	for i <= j {
		if power >= a[i] {
			power -= a[i]
			i++
			points++
			res = max(res, points)
		} else if points > 0 {
			points--
			power += a[j]
			j--
		} else {
			break
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
