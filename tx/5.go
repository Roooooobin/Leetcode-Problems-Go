package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sum := 0
	for _, num := range a {
		sum += num
	}
	//for x := -sum; x <= 0; x++ {
	//	y := x + sum
	//
	//}
}
