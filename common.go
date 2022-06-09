package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func init() {
	MOD := int(1e9 + 7)
	directions := [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	fmt.Println(MOD, directions)
}
