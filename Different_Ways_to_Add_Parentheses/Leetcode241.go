package main

import "strconv"

func diffWaysToCompute(expression string) []int {

	isDigit := func(input string) bool {
		_, err := strconv.Atoi(input)
		if err != nil {
			return false
		}
		return true
	}
	if isDigit(expression) {
		num, _ := strconv.Atoi(expression)
		return []int{num}
	}
	res := make([]int, 0)
	for i, c := range expression {
		// operator
		if c < '0' || c > '9' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, l := range left {
				for _, r := range right {
					if c == '*' {
						res = append(res, l*r)
					} else if c == '+' {
						res = append(res, l+r)
					} else {
						res = append(res, l-r)
					}
				}
			}
		}
	}
	return res
}

/*
https://leetcode.cn/problems/different-ways-to-add-parentheses/solution/pythongolang-fen-zhi-suan-fa-by-jalan/
- -分治
本质上就是遍历i,算出左右两边可能的结果然后组合
*/
