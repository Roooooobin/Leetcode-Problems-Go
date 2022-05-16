package main

import "sort"

func isAlienSorted(words []string, order string) bool {

	n := len(words)
	sorted := make([]string, n)
	for i := 0; i < n; i++ {
		sorted[i] = words[i]
	}
	alphaOrder := make([]int, 26)
	for i, c := range order {
		alphaOrder[c-'a'] = i
	}
	sort.Slice(sorted, func(i, j int) bool {
		a, b := sorted[i], sorted[j]
		x, y := 0, 0
		for x < len(a) && y < len(b) {
			if a[x] != b[y] {
				return alphaOrder[a[x]-'a'] < alphaOrder[b[y]-'a']
			}
			x++
			y++
		}
		if x == len(a) {
			return true
		}
		return false
	})
	for i := 0; i < n; i++ {
		if sorted[i] != words[i] {
			return false
		}
	}
	return true
}
