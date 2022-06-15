package main

import "sort"

func longestStrChain(words []string) int {

	n := len(words)
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	dp := make(map[string]int)
	res := 1
	for i := 1; i < n; i++ {
		w := words[i]
		for j := 0; j < len(w); j++ {
			dp[w] = max(dp[w], dp[w[:j]+w[j+1:]]+1)
		}
		res = max(res, dp[w])
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
