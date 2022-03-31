package main

func missingRolls(rolls []int, mean int, n int) []int {

	m := len(rolls)
	mSum := 0
	for _, roll := range rolls {
		mSum += roll
	}
	nSum := mean*(m+n) - mSum
	if nSum < n || nSum > n*6 {
		return []int{}
	}
	t := nSum / n
	left := nSum % n
	res := make([]int, n)
	for i := 0; i < n-left; i++ {
		res[i] = t
	}
	for i := n - left; i < n; i++ {
		res[i] = t + 1
	}
	return res
}
