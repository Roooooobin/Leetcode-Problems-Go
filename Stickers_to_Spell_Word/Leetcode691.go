package main

func minStickers(stickers []string, target string) int {
	m := len(target)
	dp := make([]int, 1<<m)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	var dpFunc func(int) int
	dpFunc = func(mask int) int {
		if dp[mask] != -1 {
			return dp[mask]
		}
		dp[mask] = m + 1
		for _, sticker := range stickers {
			left := mask
			cnt := [26]int{}
			for _, c := range sticker {
				cnt[c-'a']++
			}
			for i, c := range target {
				if mask>>i&1 == 1 && cnt[c-'a'] > 0 {
					cnt[c-'a']--
					left ^= 1 << i
				}
			}
			if left < mask {
				dp[mask] = min(dp[mask], dpFunc(left)+1)
			}
		}
		return dp[mask]
	}
	ans := dpFunc(1<<m - 1)
	if ans <= m {
		return ans
	}
	return -1
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
