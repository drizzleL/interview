package main

import "math"

func findClosest(words []string, word1 string, word2 string) int {
	ret := math.MaxInt32
	idx1, idx2 := -1, -1
	for i, w := range words {
		if w == word1 {
			idx1 = i
		}
		if w == word2 {
			idx2 = i
		}
		if idx1 != -1 && idx2 != -1 {
			gap := idx1 - idx2
			if gap < 0 {
				gap = -gap
			}
			ret = min(ret, gap)
		}
	}
	return ret
}
