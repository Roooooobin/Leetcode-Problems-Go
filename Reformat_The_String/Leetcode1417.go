package main

import "unicode"

func reformat(s string) string {

	numbers := make([]rune, 0)
	letters := make([]rune, 0)
	for _, c := range s {
		if unicode.IsDigit(c) {
			numbers = append(numbers, c)
		} else {
			letters = append(letters, c)
		}
	}
	//if len(numbers) != len(letters)+1 && len(letters) != len(numbers)+1 && len(numbers) != len(letters) {
	//	return ""
	//}
	if abs(len(numbers)-len(letters)) > 1 {
		return ""
	}
	if len(letters) == len(numbers)+1 {
		letters, numbers = numbers, letters
	}
	res := make([]rune, 0)
	for i := 0; i < len(s); i++ {
		if i&1 == 0 {
			res = append(res, numbers[i/2])
		} else {
			res = append(res, letters[i/2])
		}
	}
	return string(res)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
