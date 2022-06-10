package main

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}
	hash := make([]int, 26)
	for _, c := range s {
		hash[c-'a']++
	}
	minm := 0
	for i := range hash {
		if hash[i]&1 == 1 {
			minm++
		}
	}
	return minm <= k
}
