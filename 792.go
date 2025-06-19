package main

import "sort"

func numMatchingSubseq(s string, words []string) int {
	var dict [26][]int
	for i, c := range s {
		dict[c-'a'] = append(dict[c-'a'], i)
	}
	check := func(w string) bool {
		var idx int
		for _, c := range w {
			d := int(c - 'a')
			idx2 := sort.SearchInts(dict[d], idx)
			if idx2 == len(dict[d]) {
				return false
			}
			idx = dict[d][idx2] + 1
		}
		return true
	}
	var ret int
	for _, w := range words {
		if check(w) {
			ret += 1
		}
	}
	return ret
}
