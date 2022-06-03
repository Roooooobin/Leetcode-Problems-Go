package main

import "strings"

func addSpaces(s string, spaces []int) string {

	res := make([]string, 0)
	for i, space := range spaces {
		if i == 0 {
			res = append(res, s[:space])
		} else {
			res = append(res, s[spaces[i-1]:space])
		}
		if i == len(spaces)-1 {
			res = append(res, s[space:])
		}
	}
	return strings.Join(res, " ")
}
