package main

import (
	"fmt"
)

func removeOccurrences(s string, part string) string {

	remove := func(cur string) (string, bool) {
		for i := 0; i <= len(cur)-len(part); i++ {
			idx := 0
			for j := i; idx < len(part) && cur[j] == part[idx]; j++ {
				idx++
			}
			if idx == len(part) {
				return cur[:i] + cur[i+len(part):], true
			}
		}
		return cur, false
	}
	cur := s
	flag := true
	for true {
		cur, flag = remove(cur)
		if !flag {
			return cur
		}
	}
	return ""
}

func main() {
	fmt.Println(removeOccurrences("daabcbaabcbc", "abc"))
}
