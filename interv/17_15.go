package main

import (
	"sort"
	"strings"
)

func longestWord(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return strings.Compare(words[i], words[j]) < 0
		}
		return len(words[i]) > len(words[j])
	})
	dict := map[string]bool{}
	for _, w := range words {
		dict[w] = true
	}
	var helper func(w string) bool
	helper = func(w string) bool {
		if len(w) == 0 {
			return true
		}
		for size := 1; size <= len(w); size++ {
			if dict[w[:size]] && helper(w[size:]) {
				return true
			}
		}
		return false
	}
	for _, w := range words {
		delete(dict, w)
		if helper(w) {
			return w
		}
	}
	return ""
}
