package main

func minSteps(s string, t string) int {

	hash1 := make([]int, 26)
	hash2 := make([]int, 26)
	for _, c := range s {
		hash1[c-'a']++
	}
	for _, c := range t {
		hash2[c-'a']++
	}
	res := 0
	for i := 0; i < 26; i++ {
		if hash1[i] > hash2[i] {
			res += hash1[i] - hash2[i]
		} else if hash1[i] < hash2[i] {
			res += hash2[i] - hash1[i]
		}
	}
	return res
}
