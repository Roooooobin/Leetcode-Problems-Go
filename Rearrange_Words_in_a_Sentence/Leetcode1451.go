package main

import (
	"sort"
	"strings"
)

func arrangeWords(text string) string {

	ss := strings.Split(text, " ")
	first := []byte(ss[0])
	first[0] += 32
	ss[0] = string(first)
	sort.SliceStable(ss, func(i, j int) bool {
		si, sj := ss[i], ss[j]
		return len(si) < len(sj)
	})
	first = []byte(ss[0])
	first[0] -= 32
	ss[0] = string(first)
	return strings.Join(ss, " ")
}
