package main

func minDeletions(s string) int {

	count := make([]int, 26)
	for _, c := range s {
		count[c-'a']++
	}
	mp := make(map[int]int)
	for _, cnt := range count {
		v := mp[cnt]
		mp[cnt] = v + 1
	}
	res := 0
	for _, cnt := range count {
		if mp[cnt] > 1 {
			mp[cnt]--
			for cnt > 0 && mp[cnt] != 0 {
				res++
				cnt--
			}
			mp[cnt] = 1
		}
	}
	return res
}
