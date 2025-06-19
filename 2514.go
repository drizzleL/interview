package main

import "strings"

func countAnagrams(s string) int {
	ret := 1
	for _, str := range strings.Fields(s) {
		var dict [26]int
		for _, c := range str {
			dict[c-'a'] += 1
		}
		size := len(str)
		for _, v := range dict {
			ret *= combination(size, v)
			ret %= 1e9 + 7
			size -= v
		}
	}
	return ret
}
