package main

func maximalRectangle(matrix [][]byte) (res int) {

	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	heights := make([]int, n)
	// 每一列计算高度
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		res = max(res, largestRectangleArea(heights))
	}
	return
}

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

/*
https://leetcode.cn/problems/maximal-rectangle/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-1-8/
- -单调栈
84的应用
*/
