package main

func lengthLongestPath(input string) int {

	res := 0
	n := len(input)
	level := make([]int, n+1)
	for i := 0; i < n; {
		depth := 1
		for ; i < n && input[i] == '\t'; i++ {
			depth++
		}
		length, isFile := 0, false
		for ; i < n && input[i] != '\n'; i++ {
			if input[i] == '.' {
				isFile = true
			}
			length++
		}
		i++

		if depth > 1 {
			length += level[depth-1] + 1
		}
		if isFile {
			res = max(res, length)
		} else {
			level[depth] = length
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
