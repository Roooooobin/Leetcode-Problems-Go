package main

func balancedString(s string) int {

	cnt := make([]int, 128)
	for _, c := range s {
		cnt[c]++
	}
	target := len(s) / 4
	check := func() bool {
		if cnt['Q'] > target || cnt['W'] > target || cnt['E'] > target || cnt['R'] > target {
			return false
		}
		return true
	}
	if check() {
		return 0
	}

	res := len(s)
	r := 0
	for l, c := range s {
		for r < len(s) && !check() {
			cnt[s[r]]--
			r++
		}
		if !check() {
			break
		}
		res = min(res, r-l)
		cnt[c]++
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
- -滑动窗口
*/
