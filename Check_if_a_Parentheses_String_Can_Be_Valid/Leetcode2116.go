package main

func canBeValid(s string, locked string) bool {

	if len(s)%2 == 1 {
		return false
	}
	left, right, all := 0, 0, 0
	n := len(s)
	for i := 0; i < n; i++ {
		if locked[i] == '0' {
			all++
		} else {
			if s[i] == '(' {
				left++
			} else {
				right++
			}
		}
		if right > left+all {
			return false
		}
	}
	left, right, all = 0, 0, 0
	for i := n - 1; i >= 0; i-- {
		if locked[i] == '0' {
			all++
		} else {
			if s[i] == '(' {
				left++
			} else {
				right++
			}
		}
		if left > right+all {
			return false
		}
	}
	return true
}
