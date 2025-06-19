package main

import (
	"strings"
)

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	if len(sentence1) < len(sentence2) {
		sentence1, sentence2 = sentence2, sentence1
	}
	s1, s2 := strings.Split(sentence1, " "), strings.Split(sentence2, " ")
	i, j := 0, len(s2)-1
	for ; i < len(s2); i++ {
		if s2[i] != s1[i] {
			break
		}
	}
	for m := len(s1) - 1; j >= i; j, m = j-1, m-1 {
		if s2[j] != s1[m] {
			break
		}
	}
	return i > j
}
