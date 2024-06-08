package main

import "sort"

func minimumLengthEncoding(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	var ret int
	dict := map[string]bool{}
	for _, word := range words {
		if dict[word] {
			continue
		}
		ret += len(word) + 1
		for i := len(word) - 1; i >= 0; i-- {
			dict[word[i:]] = true
		}
	}
	return ret
}
