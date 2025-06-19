package main

import "strings"

func isCircularSentence(sentence string) bool {
	strs := strings.Fields(sentence)
	for i := 1; i < len(strs); i++ {
		if strs[i][0] != strs[i-1][len(strs[i-1])-1] {
			return false
		}
	}
	return strs[0][0] == strs[0][len(strs[0])-1]
}
