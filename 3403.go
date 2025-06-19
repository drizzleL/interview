package main

import (
	"strings"
)

func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	var ret string
	for i := 0; i < len(word); i++ {
		s := word[i:min(len(word), i+len(word)-numFriends+1)]
		if strings.Compare(s, ret) > 0 {
			ret = s
		}
	}
	return ret
}
