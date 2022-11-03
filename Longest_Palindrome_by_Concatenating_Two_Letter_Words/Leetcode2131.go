package main

func longestPalindrome(words []string) (res int) {

	counter := make(map[string]int)
	for _, word := range words {
		counter[word]++
	}
	midFlag := false
	for word, count := range counter {
		reverse := string([]byte{word[1], word[0]})
		if word == reverse {
			if !midFlag && count&1 == 1 {
				midFlag = true
			}
			res += 2 * (count / 2 * 2)
			// avoid redundancy
		} else if word > reverse {
			res += 4 * min(counter[word], counter[reverse])
		}
	}
	if midFlag {
		res += 2
	}
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
