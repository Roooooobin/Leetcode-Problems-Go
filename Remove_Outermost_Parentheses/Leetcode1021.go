package main

import "fmt"

func removeOuterParentheses(s string) string {

	n := len(s)
	stack := make([]int, 0)
	stack = append(stack, 0)
	res := ""
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if len(stack) == 1 && i-stack[0] > 1 {
				res += s[stack[0]+1 : i]
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, i)
		}
	}
	return res
}

func removeOuterParentheses2(s string) string {
	var ans, st []rune
	for _, c := range s {
		if c == ')' {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans = append(ans, c)
		}
		if c == '(' {
			st = append(st, c)
		}
	}
	return string(ans)
}

//
//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/remove-outermost-parentheses/solution/shan-chu-zui-wai-ceng-de-gua-hao-by-leet-sux0/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	fmt.Println(removeOuterParentheses("()()()()(())"))
}
