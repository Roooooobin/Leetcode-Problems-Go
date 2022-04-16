package main

import "unicode"

func mostCommonWord(paragraph string, banned []string) string {

	n := len(paragraph)
	banSet := make(map[string]struct{})
	for _, s := range banned {
		banSet[s] = struct{}{}
	}
	freq := make(map[string]int)
	maxFreq := 0
	var word []byte
	for i := 0; i <= n; i++ {
		if i < n && unicode.IsLetter(rune(paragraph[i])) {
			word = append(word, byte(unicode.ToLower(rune(paragraph[i]))))
		} else if word != nil {
			s := string(word)
			if _, ok := banSet[s]; !ok {
				freq[s]++
				maxFreq = max(maxFreq, freq[s])
			}
			word = nil
		}
	}
	for s, f := range freq {
		if f == maxFreq {
			return s
		}
	}
	return ""
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
