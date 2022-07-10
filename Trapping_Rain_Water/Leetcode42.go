package main

func trap(a []int) int {
	res := 0
	stack := make([]int, 0)
	for i := 0; i < len(a); i++ {
		for len(stack) > 0 && a[stack[len(stack)-1]] < a[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			res += (min(a[i], a[left]) - a[top]) * (i - left - 1)
		}
		stack = append(stack, i)
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
- -单调栈

*/
