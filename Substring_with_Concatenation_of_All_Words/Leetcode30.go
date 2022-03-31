package main

func findSubstring(s string, words []string) []int {

	m := len(words[0])
	n := len(words)
	totLen := m * n
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}
	res := make([]int, 0)
	for i := 0; i <= len(s)-totLen; i++ {
		tmpMap := make(map[string]int)
		flag := true
		for j := i; j <= i+totLen-m; j += m {
			word := s[j : j+m]
			if _, ok := wordMap[word]; !ok {
				flag = false
				break
			}
			tmpMap[word]++
			if tmpMap[word] > wordMap[word] {
				flag = false
				break
			}
			if !flag {
				break
			}
		}
		if flag {
			res = append(res, i)
		}
	}
	return res
}
