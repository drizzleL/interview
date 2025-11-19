package main

import "strings"

func canBeTypedWords(text string, brokenLetters string) int {
	var dict [26]bool
	for _, c := range brokenLetters {
		dict[c-'a'] = true
	}
	var ret int
	for _, word := range strings.Fields(text) {
		var flag bool
		for _, c := range word {
			if dict[c-'a'] {
				flag = true
				break
			}
		}
		if flag {
			continue
		}
		ret += 1
	}
	return ret
}
