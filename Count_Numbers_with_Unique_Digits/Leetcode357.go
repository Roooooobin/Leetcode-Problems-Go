package main

func countNumbersWithUniqueDigits(n int) int {

	res := 0
	if n == 0 {
		return 1
	} else if n == 1 {
		return 9
	}
	if n >= 2 {
		res += 91
	}
	for k := 3; k <= n; k++ {
		a := calcCombinationNum(10, k-1)
		b := calcCombinationNum(9, k-1)
		res += (a-b)*(10-k) + b*(9-k)
	}
	return res
}

func calcCombinationNum(n, c int) int {

	product := 1
	for i := n; i >= n-c+1; i-- {
		product *= i
	}
	return product
}
