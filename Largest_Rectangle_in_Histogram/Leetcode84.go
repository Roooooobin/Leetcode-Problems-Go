package main

func largestRectangleArea(heights []int) int {

	stack := make([]int, 0)
	heights = append(heights, 0)
	res := 0
	i, n := 0, len(heights)
	for i < n {
		if len(stack) == 0 || heights[stack[len(stack)-1]] <= heights[i] {
			stack = append(stack, i)
			i++
		} else {
			topIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			preIdx := -1
			if len(stack) > 0 {
				preIdx = stack[len(stack)-1]
			}
			res = max(res, heights[topIdx]*(i-preIdx-1))
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

func main() {
	largestRectangleArea([]int{2, 1, 5, 6, 2, 3, 3, 3})
}

/*
单调栈
保持一个单调递增的栈,如果当前值比栈顶小,那么今后的都将和当前值之前的所有值隔断,
可以把当前值之前的都算好,栈顶值 * (i-preIdx-1),注意区间是i-栈顶前的那个坐标(不太好解释,做个图更容易想到)
*/
