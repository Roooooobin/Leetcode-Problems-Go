package main

import (
	"strings"
)

func reorderSpaces(text string) string {

	words := strings.Fields(text)
	spaceCount := strings.Count(text, " ")
	wordCount := len(words) - 1
	if wordCount == 0 {
		return words[0] + strings.Repeat(" ", spaceCount)
	}
	return strings.Join(words, strings.Repeat(" ", spaceCount/wordCount)) + strings.Repeat(" ", spaceCount%wordCount)
}
