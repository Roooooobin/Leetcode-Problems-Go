package main

func maxLengthBetweenEqualCharacters(s string) int {

	res := -1
	first := [26]int{}
	for i := range first {
		first[i] = -1
	}
	for i, c := range s {
		c -= 'a'
		if first[c] == -1 {
			first[c] = i
		} else {
			res = max(res, i-first[c]-1)
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
- -模拟
记录每个字母出现的第一个位置
*/
