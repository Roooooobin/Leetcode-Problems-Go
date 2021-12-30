package main

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	x := 0
	for n := 1; n <= k; n++ {
		x = (x*10 + 1) % k
		if x%k == 0 {
			return n
		}
	}
	return -1
}
