package main

import (
	"strings"
)

func largestWordCount(messages []string, senders []string) string {

	mp := make(map[string]int)
	maxm := 0
	res := ""
	for i, sender := range senders {
		v := mp[sender]
		v += len(strings.Split(messages[i], " "))
		mp[sender] = v
		if maxm < v {
			maxm = v
			res = sender
		} else if maxm == v && res < sender {
			res = sender
		}
	}
	return res
}
