package main

func concatenatedBinary(n int) (res int) {

	MOD := int(1e9 + 7)
	shift := 0
	for i := 1; i < n+1; i++ {
		if (i & (i - 1)) == 0 {
			shift++
		}
		res = ((res << shift) + i) % MOD
	}
	return res
}

/*
https://leetcode.cn/problems/concatenation-of-consecutive-binary-numbers/solution/lian-jie-lian-xu-er-jin-zhi-shu-zi-by-ze-t40j/
*/
