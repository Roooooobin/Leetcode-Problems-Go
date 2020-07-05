package main

func hammingDistance(x int, y int) int {
	z := x ^ y
	var dist int
	for z != 0 {
		dist++
		z = z & (z - 1)
	}
	return dist
}
