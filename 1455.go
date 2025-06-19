package main

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	for i, s := range strings.Fields(sentence) {
		if len(s) < len(searchWord) {
			continue
		}
		if s[:len(searchWord)] == searchWord {
			return i + 1
		}
	}
	return -1
}
