package main

func consecutiveNumbersSum(n int) int {

	check := func(k int) bool {
		if k&1 == 1 {
			return n%k == 0
		}
		return n%k != 0 && 2*n%k == 0
	}
	res := 0
	for k := 1; k*(k+1) <= n*2; k++ {
		if check(k) {
			res++
		}
	}
	return res
}
