package main

func longestSubsequence(s string, k int) int {

	res := 0
	for _, c := range s {
		if c == '0' {
			res++
		}
	}
	cost := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '1' {
			if len(s)-i-1 >= 32 {
				break
			}
			cost += 1 << (len(s) - i - 1)
		}
		if cost <= k {
			res += int(s[i] - '0')
		} else {
			break
		}
	}
	return res
}

/*
https://leetcode.com/problems/longest-binary-subsequence-less-than-or-equal-to-k/discuss/2168227/JavaC%2B%2BPython-One-Pass-O(n)
- -贪心
拿所有0,从后往前尽可能拿1
*/
