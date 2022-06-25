package main

import "math/rand"

type Solution struct {
	mp    map[int]int
	bound int
}

func Constructor(n int, blacklist []int) Solution {
	m := len(blacklist)
	bound := n - m
	black := make(map[int]bool)
	for _, x := range blacklist {
		if x >= bound {
			black[x] = true
		}
	}
	mp := make(map[int]int, m-len(black))
	b := bound
	for _, x := range blacklist {
		if x < bound {
			for black[b] {
				b++
			}
			mp[x] = b
			b++
		}
	}
	return Solution{mp, bound}
}

func (this *Solution) Pick() int {
	x := rand.Intn(this.bound)
	if this.mp[x] > 0 {
		return this.mp[x]
	}
	return x
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */

//https://leetcode.cn/problems/random-pick-with-blacklist/solution/hei-ming-dan-zhong-de-sui-ji-shu-by-leet-cyrx/
