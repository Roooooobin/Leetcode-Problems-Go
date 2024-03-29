package main

import (
	"strconv"
)

func splitString(s string) bool {

	n := len(s)
	var backtracking func(idx, pre int)
	res := false
	backtracking = func(idx, pre int) {
		if res {
			return
		}
		if idx == n {
			res = true
		}
		for i := idx; i < n; i++ {
			num, err := strconv.Atoi(s[idx : i+1])
			if err != nil {
				continue
			}
			if num == pre-1 {
				backtracking(i+1, num)
			}
		}
	}
	for i := 0; i < n-1; i++ {
		num, err := strconv.Atoi(s[:i+1])
		if err == nil {
			backtracking(i+1, num)
		}
	}
	return res
}

/*
- -回溯
记录一个pre, 然后回溯确认当前切分的num是否为pre-1
*/
