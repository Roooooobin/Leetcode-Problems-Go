package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	res := make([]int, 0)
	l, r := 0, n-1
	round := 0
	for l <= r {
		if round&1 == 0 {
			if a[l] < a[r] {
				res = append(res, a[r])
				r--
			} else {
				res = append(res, a[l])
				l++
			}
		} else {
			if a[l] > a[r] {
				res = append(res, a[r])
				r--
			} else {
				res = append(res, a[l])
				l++
			}
		}
		round++
	}
	for i := 0; i < n; i++ {
		fmt.Print(res[i])
		fmt.Print(" ")
	}
	fmt.Println()
}
