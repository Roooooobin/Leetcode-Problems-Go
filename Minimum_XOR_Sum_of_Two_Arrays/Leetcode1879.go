package main

import (
	"fmt"
	"math/bits"
)

func minimumXORSum(nums1 []int, nums2 []int) int {

	n := len(nums2)
	m := 1 << n
	dp := make([]int, m)
	for i := 1; i < m; i++ {
		dp[i] = 0x3f3f3f3f
	}
	dp[0] = 0
	for i := range dp {
		c := bits.OnesCount(uint(i))
		for j := 0; j < n; j++ {
			if i&(1<<j) == 0 {
				s := i | 1<<j
				// 位运算优先级低于加减！！！！！！
				dp[s] = min(dp[s], dp[i]+(nums1[j]^nums2[c]))
			}
		}
	}
	return dp[m-1]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	fmt.Println(minimumXORSum([]int{1, 0, 3}, []int{5, 3, 4}))
}
