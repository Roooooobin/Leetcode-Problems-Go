package main

import "strconv"

func digitSum(s string, K int) string {

	calc := func(s string) string {
		sum := 0
		for _, c := range s {
			sum += int(c - '0')
		}
		return strconv.Itoa(sum)
	}
	for len(s) > K {
		next := ""
		for k := 0; k < len(s); k += K {
			end := k + K
			if end > len(s) {
				end = len(s)
			}
			next += calc(s[k:end])
		}
		s = next
	}
	return s
}
