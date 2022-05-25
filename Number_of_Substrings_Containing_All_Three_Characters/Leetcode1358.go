package main

func numberOfSubstrings(s string) int {

	res := 0
	n := len(s)
	l := 0
	abc := make([]int, 3)
	for r := 0; r < n; r++ {
		abc[s[r]-'a']++
		for l < r && satisfy(abc) {
			res += n - r
			abc[s[l]-'a']--
			l++
		}
	}
	return res
}

func satisfy(abc []int) bool {
	return abc[0] > 0 && abc[1] > 0 && abc[2] > 0
}
