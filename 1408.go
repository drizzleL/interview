package main

import "strings"

func stringMatching(words []string) []string {
	var ret []string
	for i := 0; i < len(words); i++ {
		check := func(w1, w2 string) bool {
			return strings.Contains(w2, w1)
		}
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			if check(words[i], words[j]) {
				ret = append(ret, words[i])
			}
		}
	}
	return ret
}
