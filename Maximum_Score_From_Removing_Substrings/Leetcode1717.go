package main

func maximumGain(s string, x int, y int) int {

	res := 0
	stack := make([]byte, 0)
	val := [2]int{x, y}
	target := "ab"
	if x < y {
		target = "ba"
	}
	for i := range s {
		c := s[i]
		if c == target[1] {
			if len(stack) > 0 && stack[len(stack)-1] == target[0] {
				res += val[target[0]-'a']
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, c)
			}
		} else {
			stack = append(stack, c)
		}
	}
	s = string(stack)
	stack = make([]byte, 0)
	for i := range s {
		c := s[i]
		if c == target[0] {
			if len(stack) > 0 && stack[len(stack)-1] == target[1] {
				res += val[target[1]-'a']
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, c)
			}
		} else {
			stack = append(stack, c)
		}
	}
	return res
}

/*
- -贪心+栈
谁(ab/ba)收益大先算谁
*/
