package main

func characterReplacement(s string, k int) int {

	slidingWindow := func(c byte) (res int) {
		cnt := 0
		n := len(s)
		for l, r := 0, 0; r < n; r++ {
			if s[r] != c {
				cnt++
			}
			for cnt > k {
				if s[l] != c {
					cnt--
				}
				l++
			}
			if res < r-l+1 {
				res = r - l + 1
			}
		}
		return
	}
	res := 0
	for c := byte('A'); c <= 'Z'; c++ {
		r := slidingWindow(c)
		if res < r {
			res = r
		}
	}
	return res
}
