package main

import "sort"

func maxPalindromesAfterOperations(words []string) int {
	cnts := make([]int, 26)
	for _, w := range words {
		for _, c := range w {
			cnts[c-'a'] += 1
		}
	}
	var pairs int
	for _, v := range cnts {
		pairs += v / 2
	}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	for i, w := range words {
		size := len(w)
		pairs -= size / 2
		if pairs < 0 {
			return i
		}
	}
	return len(words)
}
