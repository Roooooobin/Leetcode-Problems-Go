package main

import (
	"strings"
)

func removeKdigits(num string, k int) string {

	stack := make([]byte, 0)
	for i := 0; i < len(num); i++ {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			k--
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num[i])
	}
	if k > 0 {
		stack = stack[:len(stack)-k]
	}
	res := strings.TrimLeft(string(stack), "0")
	if res == "" {
		res = "0"
	}
	return res
}
