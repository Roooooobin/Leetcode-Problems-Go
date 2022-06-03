package main

import "math/bits"

func maximumANDSum(nums []int, numSlots int) int {

	res := 0
	dp := make([]int, 1<<(2*numSlots))
	for i := range dp {
		c := bits.OnesCount(uint(i))
		if c >= len(nums) {
			continue
		}
		for j := 0; j < 2*numSlots; j++ {
			if i>>j&1 == 0 {
				s := i | 1<<j
				dp[s] = max(dp[s], dp[i]+(j/2+1)&nums[c])
				res = max(res, dp[s])
			}
		}
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

//作者：endlesscheng
//链接：https://leetcode.cn/problems/maximum-and-sum-of-array/solution/zhuang-tai-ya-suo-dp-by-endlesscheng-5eqn/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
