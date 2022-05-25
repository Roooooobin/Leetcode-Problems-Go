package main

func minimumDeletions(s string) int {

	n := len(s)
	count := make([]int, n)
	b := 0
	a := 0
	for i := 0; i < n; i++ {
		count[i] = b
		if s[i] == 'b' {
			b++
		}
	}
	for i := n - 1; i >= 0; i-- {
		count[i] += a
		if s[i] == 'a' {
			a++
		}
	}
	if a == 0 || b == 0 {
		return 0
	}
	res := n
	for i := 0; i < n; i++ {
		res = min(res, count[i])
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
