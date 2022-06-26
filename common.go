package main

import (
	"fmt"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func init() {
	MOD := int(1e9 + 7)
	directions := [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	fmt.Println(MOD, directions)
}

//
//type maxHeap []int
//
//func (h maxHeap) Len() int            { return len(h) }
//func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
//func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
//func (h *maxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
//func (h *maxHeap) Pop() interface{}   { a := *h; n := len(*h); v := a[n-1]; *h = a[:n-1]; return v }

func main() {
	convert := func(s string) string {
		return strings.Join(strings.Split(s, " "), "_")
	}
	fmt.Println(convert(""))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
