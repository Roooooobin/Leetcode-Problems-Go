package main

import "strings"

func maxRepeating(sequence string, word string) int {

	tmp := word
	res := 0
	for strings.Contains(sequence, tmp) {
		tmp = tmp + word
		res++
	}
	return res
}
