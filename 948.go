package main

import (
	"sort"
)

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	l, r := 0, len(tokens)-1
	var ret, curr int
	for l <= r {
		for l <= r && power >= tokens[l] {
			power -= tokens[l]
			curr++
			l++
		}
		ret = max(ret, curr)
		if l >= r || curr < 1 {
			break
		}
		power += tokens[r]
		r--
		curr -= 1
	}
	return ret
}
