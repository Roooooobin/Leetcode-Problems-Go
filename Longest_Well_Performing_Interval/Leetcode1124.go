package main

func longestWPI(hours []int) int {

	preSum := make([]int, len(hours)+1)
	preSum[0] = 0
	for i := 0; i < len(hours); i++ {
		if hours[i] > 8 {
			hours[i] = 1
		} else {
			hours[i] = -1
		}
		preSum[i+1] = preSum[i] + hours[i]
	}

	stack := make([]int, 0)
	for i, sum := range preSum {
		if len(stack) == 0 || preSum[stack[len(stack)-1]] > sum {
			stack = append(stack, i)
		}
	}
	res := 0
	for i := len(preSum) - 1; i >= 0; i-- {
		for len(stack) != 0 && preSum[stack[len(stack)-1]] < preSum[i] {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = max(res, i-cur)
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
- -单调栈(难)
*/
