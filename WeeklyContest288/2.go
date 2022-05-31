package main

import "fmt"

func minimizeResult(expression string) string {

	idxOfPlus := 0
	for i, c := range expression {
		if c == '+' {
			idxOfPlus = i
			break
		}
	}
	calc := func(s string) int {
		num1 := 0
		i := 0
		n := len(s)
		for i < n {
			if s[i] == '(' {
				break
			}
			num1 = num1*10 + int(s[i]-'0')
			i++
		}
		if num1 == 0 {
			num1 = 1
		}
		i++
		num2 := 0
		for i < n {
			if s[i] == '+' {
				break
			}
			num2 = num2*10 + int(s[i]-'0')
			i++
		}
		i++
		num3 := 0
		for i < n {
			if s[i] == ')' {
				break
			}
			num3 = num3*10 + int(s[i]-'0')
			i++
		}
		i++
		num4 := 0
		for i < n {
			num4 = num4*10 + int(s[i]-'0')
			i++
		}
		if num4 == 0 {
			num4 = 1
		}
		return num1 * (num2 + num3) * num4
	}
	res := -1
	resS := ""
	for i := 0; i < idxOfPlus; i++ {
		newS := ""
		for j := idxOfPlus + 2; j <= len(expression); j++ {
			if i == 0 {
				newS = "(" + expression[i:j]
			} else {
				newS = expression[:i] + "(" + expression[i:j]
			}
			if j == len(expression) {
				newS += ")"
			} else {
				newS += ")" + expression[j:]
			}
			cur := calc(newS)
			if res == -1 || res > cur {
				res = cur
				resS = newS
			}
		}
	}
	return resS
}

func main() {
	fmt.Println(minimizeResult("999+999"))
}
