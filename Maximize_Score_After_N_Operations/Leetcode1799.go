package main

import "math/bits"

/*
https://leetcode.cn/problems/maximize-score-after-n-operations/solution/1799-n-ci-cao-zuo-hou-de-zui-da-fen-shu-tagge/
*/
func maxScore(nums []int) int {

	n := len(nums)
	dp := make([]int, 1<<n)
	for i := range dp {
		c := bits.OnesCount(uint(i))
		if c&1 == 1 {
			continue
		}
		for j := 0; j < n; j++ {
			if i&(1<<j) > 0 {
				for k := j + 1; k < n; k++ {
					if i&(1<<k) > 0 {
						dp[i] = max(dp[i], dp[i^(1<<k)^(1<<j)]+c/2*gcd(nums[j], nums[k]))
					}
				}
			}
		}
	}
	return dp[(1<<n)-1]
}

func gcd(x, y int) int {

	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
