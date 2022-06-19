package main

import (
	"sort"
	"strings"
)

func minimumLengthEncodingWA(words []string) int {

	res := 0
	count := 0
	for i, word := range words {
		flag := false
		for j, word2 := range words {
			if i == j || len(word) > len(word2) {
				continue
			}
			target := len(word2) - len(word)
			if strings.LastIndex(word2, word) == target {
				flag = true
				break
			}
		}
		if !flag {
			res += len(word)
			count++
		}
	}
	if count == 0 {
		return len(words[0]) + 1
	}
	return res + count
}

func reverse(s []byte) string {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
	return string(s)
}

func minimumLengthEncoding(words []string) int {

	t := Trie{}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	res := 0
	count := 0
	for _, word := range words {
		reversedWord := reverse([]byte(word))
		if !t.IsPrefix(reversedWord) {
			res += len(word)
			count++
		}
		t.Insert(reversedWord)
	}
	return res + count
}

type Trie struct {
	nodes [26]*Trie
}

func (t *Trie) Insert(s string) {
	cur := t
	for _, c := range s {
		c -= 'a'
		if cur.nodes[c] == nil {
			cur.nodes[c] = &Trie{}
		}
		cur = cur.nodes[c]
	}
}

func (t *Trie) IsPrefix(prefix string) bool {
	cur := t
	for _, c := range prefix {
		c -= 'a'
		if cur.nodes[c] == nil {
			return false
		}
		cur = cur.nodes[c]
	}
	return true
}
