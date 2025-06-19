package main

import "strings"

func prefixCount(words []string, pref string) int {
	var ret int
	for _, w := range words {
		if strings.HasPrefix(w, pref) {
			ret += 1
		}
	}
	return ret
}
