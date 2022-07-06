package main

import (
	"fmt"
	"math"
	"strconv"
)

func splitIntoFibonacci(num string) []int {

	res := make([]int, 0)
	flag := false
	n := len(num)
	var backtracking func(idx int, cur []int)
	backtracking = func(idx int, cur []int) {
		if flag {
			return
		}
		if idx == n && len(cur) > 2 {
			flag = true
			res = append([]int{}, cur...)
			return
		} else if idx == n {
			return
		}
		for i := idx; i < n; i++ {
			x, _ := strconv.Atoi(num[idx : i+1])
			// 关键, 不能超过int上限
			if x >= math.MaxInt32 {
				break
			}
			if len(cur) < 2 || x-cur[len(cur)-1] == cur[len(cur)-2] {
				cur = append(cur, x)
				backtracking(i+1, cur)
				cur = cur[:len(cur)-1]
			}
			if num[idx] == '0' {
				break
			}
		}
	}
	for i := range num {
		prev, _ := strconv.Atoi(num[:i+1])
		backtracking(i+1, []int{prev})
		if prev == 0 {
			break
		}
	}
	return res
}

/*
- -回溯+剪枝
以零开头只能取0, 剪枝
超过int32!!!要直接剪枝,总是被坑
*/

func main() {
	fmt.Println(splitIntoFibonacci("1101111"))
	fmt.Println(splitIntoFibonacci("11235813"))
}
