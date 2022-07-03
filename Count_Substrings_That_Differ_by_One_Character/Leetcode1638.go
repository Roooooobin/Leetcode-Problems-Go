package main

func countSubstrings(s string, t string) int {

	m, n := len(s), len(t)
	ans := 0
	helper := func(i, j int) int {
		pre, cur := 0, 0
		res := 0
		for i < m && j < n {
			cur++
			if s[i] != t[j] {
				pre = cur
				cur = 0
			}
			res += pre
			i++
			j++
		}
		return res
	}
	for i := 0; i < m; i++ {
		ans += helper(i, 0)
	}
	for j := 1; j < n; j++ {
		ans += helper(0, j)
	}
	return ans
}

/*
https://leetcode.com/problems/count-substrings-that-differ-by-one-character/discuss/917985/JavaC%2B%2BPython-Time-O(nm)-Space-O(1)
- -牛逼思路
记录上一次pre和本次cur的最大连续相同字符,如果出现不相同字符,pre = cur, cur = 0, res += pre
*/
