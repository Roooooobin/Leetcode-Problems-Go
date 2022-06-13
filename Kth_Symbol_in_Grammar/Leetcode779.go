package main

func kthGrammar(n int, k int) int {

	m := 1 << (n - 1)
	if k == 1 {
		return 0
	}
	if k < m/2 {
		return kthGrammar(n-1, k)
	} else {
		c := kthGrammar(n-1, (k+1)/2)
		if k&1 == 0 {
			return c ^ 1
		} else {
			return c
		}
	}
}
