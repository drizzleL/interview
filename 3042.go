package main

import "strings"

func countPrefixSuffixPairs(words []string) int {
	var ret int
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if strings.HasPrefix(words[j], words[i]) && strings.HasSuffix(words[j], words[i]) {
				ret += 1
			}
		}
	}
	return ret
}
