package main

import "sort"

func makesquare(matchsticks []int) bool {

	n := len(matchsticks)
	sum := 0
	for _, matchstick := range matchsticks {
		sum += matchstick
	}
	if sum%4 != 0 {
		return false
	}
	target := sum / 4
	for _, matchstick := range matchsticks {
		if matchstick > target {
			return false
		}
	}
	// sort to reduce amount of searches
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})
	sticks := make([]int, 4)
	var dfs func(idx int, a []int)
	flag := false
	dfs = func(idx int, a []int) {
		if flag {
			return
		}
		if idx >= n {
			if check(a, target) {
				flag = true
			}
			return
		}
		cur := matchsticks[idx]
		for i := 0; i < 4; i++ {
			if a[i]+cur <= target {
				a[i] += cur
				dfs(idx+1, a)
				a[i] -= cur
			}
		}
	}
	dfs(0, sticks)
	return flag
}

func check(a []int, target int) bool {

	for _, x := range a {
		if x != target {
			return false
		}
	}
	return true
}
