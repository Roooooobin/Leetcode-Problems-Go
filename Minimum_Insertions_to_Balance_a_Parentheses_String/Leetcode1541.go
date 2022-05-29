package main

func minInsertions(s string) int {

	left, right := 0, 0
	res := 0
	for _, c := range s {
		if c == '(' {
			if right == 1 {
				if left > 0 {
					res++
					left--
				} else {
					res += 2
				}
				right--
			}
			left++
		} else {
			right++
			if right == 2 {
				right = 0
				if left == 0 {
					res++
				} else {
					left--
				}
			}
		}
	}
	if right == 1 {
		if left > 0 {
			res++
			left--
		} else {
			res += 2
		}
		right--
	}
	return res + 2*left
}
