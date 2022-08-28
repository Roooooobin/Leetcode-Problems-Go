package main

import "sort"

func preimageSizeFZF(K int) int {

	postfixZero := func(x int) int {
		cnt := 0
		for x > 0 {
			x /= 5
			cnt += x
		}
		return cnt
	}
	nx := func(k int) int {
		return sort.Search(5*k, func(x int) bool {
			return postfixZero(x) >= k
		})
	}
	return nx(K+1) - nx(K)
}

/*
- -公式推导+二分
https://leetcode.cn/problems/preimage-size-of-factorial-zeroes-function/solution/jie-cheng-han-shu-hou-k-ge-ling-by-leetc-n6vj/
*/
