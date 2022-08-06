package main

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {

	res := make([]int, n)
	stack := make([][2]int, 0)
	for _, log := range logs {
		info := strings.Split(log, ":")
		idx, _ := strconv.Atoi(info[0])
		time, _ := strconv.Atoi(info[2])
		if info[1] == "start" {
			if len(stack) > 0 {
				res[stack[len(stack)-1][0]] += time - stack[len(stack)-1][1]
				stack[len(stack)-1][1] = time
			}
			stack = append(stack, [2]int{idx, time})
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[top[0]] += time - top[1] + 1
			if len(stack) > 0 {
				stack[len(stack)-1][1] = time + 1
			}
		}
	}
	return res
}

/*
- -栈
模拟函数栈, 注意有递归调用
*/
