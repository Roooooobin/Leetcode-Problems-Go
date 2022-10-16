package main

import "fmt"

//a := make([]int, 0)
//a = append(a, 1)
//a = append(a, 0)
//for int64(len(a)) <= r {
//	curN := len(a)
//	for i := 0; i < curN; i++ {
//		a = append(a, a[i]^1)
//	}
//}
//res := 0
//for i := l - 1; i <= r-1; i++ {
//	if a[i] == 1 {
//		res++
//	}
//}
//fmt.Println(res)

func main() {
	var l, r int64
	fmt.Scan(&l, &r)
	rs := helper(r - 1)
	var ls int64
	if l != 1 {
		ls = helper(l - 2)
	}
	fmt.Println(rs - ls)
}

func helper(x int64) int64 {
	red := x % 4
	onesBefore := int64(x/4) * 2
	y := x - red
	cnt := 0
	for y != 0 && y != 1 {
		v2 := int64(1)
		for v2 <= y {
			v2 *= 2
		}
		cnt++
		v2 /= 2
		y -= v2
	}
	modes := []string{"1001", "0110"}
	mode := modes[cnt%2]
	for i := 0; i <= int(red); i++ {
		if mode[i] == '1' {
			onesBefore++
		}
	}
	return onesBefore
}
