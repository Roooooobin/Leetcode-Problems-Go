package main

import (
	"sort"
	"strconv"
)

func largestInteger(num int) int {

	s := strconv.Itoa(num)
	oddIdx := make([]int, 0)
	evenIdx := make([]int, 0)
	odds := make([]byte, 0)
	evens := make([]byte, 0)
	for i := range s {
		if (s[i]-'0')%2 == 0 {
			evens = append(evens, s[i])
			evenIdx = append(evenIdx, i)
		} else {
			oddIdx = append(oddIdx, i)
			odds = append(odds, s[i])
		}
	}
	sort.Slice(odds, func(i, j int) bool { return odds[i] > odds[j] })
	sort.Slice(evens, func(i, j int) bool { return evens[i] > evens[j] })
	res := make([]byte, len(s))
	for i, even := range evens {
		res[evenIdx[i]] = even
	}
	for i, odd := range odds {
		res[oddIdx[i]] = odd
	}
	resS := string(res)
	n, _ := strconv.Atoi(resS)
	return n
}
