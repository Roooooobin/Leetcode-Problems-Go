package main

import "strconv"

func rotatedDigits(n int) (res int) {

	check := [10]int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1}
	for i := 1; i <= n; i++ {
		isValid, diff := true, false
		for _, c := range strconv.Itoa(i) {
			c -= '0'
			if check[c] == -1 {
				isValid = false
			} else if check[c] == 1 {
				diff = true
			}
		}
		if isValid && diff {
			res++
		}
	}
	return
}
