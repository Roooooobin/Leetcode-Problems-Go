package main

func lengthOfLongestSubstring(s string) int {

	hash := make([]int, 256)
	res := 0
	for i, j := 0, 0; i < len(s); i++ {
		hash[s[i]]++
		for j < i && hash[s[i]] > 1 {
			hash[s[j]]--
			j++
		}
		res = max(res, i-j+1)
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
