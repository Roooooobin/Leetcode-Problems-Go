package main

import "unicode"

func decodeAtIndex(s string, k int) string {

	size := int64(0)
	for _, c := range s {
		if unicode.IsDigit(c) {
			size *= int64(c - '0')
		} else {
			size++
		}
	}
	K := int64(k)
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		K %= size
		if K == 0 && unicode.IsLetter(rune(c)) {
			return string(c)
		}
		if unicode.IsDigit(rune(c)) {
			size /= int64(c - '0')
		} else {
			size--
		}
	}
	return ""
}
