package main

import "strings"

func minRemoveToMakeValid(s string) string {

	stack := make([]int, 0)
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else if c == ')' {
			if len(stack) == 0 {
				stack = append(stack, i)
			} else if s[stack[len(stack)-1]] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, i)
			}
		}
	}
	sb := strings.Builder{}
	idx := 0
	for i, c := range s {
		if idx < len(stack) && i == stack[idx] {
			idx++
			continue
		} else {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}
