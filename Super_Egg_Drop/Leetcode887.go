package main

func superEggDrop(k int, n int) int {

	memo := make([][]int, k+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}
	var dp func(i, n int) int
	dp = func(i, n int) int {
		if i == 1 || n == 1 {
			return n
		}
		if memo[i][n] != 0 {
			return memo[i][n]
		}
		res := n
		lo, hi := 1, n
		// binary search to find valley
		for lo <= hi {
			mid := lo + (hi-lo)>>1
			lv := dp(i-1, mid-1)
			rv := dp(i, n-mid)
			res = min(res, max(lv, rv)+1)
			if lv < rv {
				lo = mid + 1
			} else if lv > rv {
				hi = mid - 1
			} else {
				break
			}
		}
		memo[i][n] = res
		return res
	}
	return dp(k, n)
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
