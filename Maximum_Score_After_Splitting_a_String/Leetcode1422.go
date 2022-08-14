package main

func maxScore(s string) int {

	cntOnes := 0
	for _, c := range s {
		if c == '1' {
			cntOnes++
		}
	}
	res := 0
	cntZeros := 0
	if s[0] == '0' {
		cntZeros++
	} else {
		cntOnes--
	}
	for i := 1; i < len(s); i++ {
		res = max(res, cntZeros+cntOnes)
		if s[i] == '0' {
			cntZeros++
		} else {
			cntOnes--
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
