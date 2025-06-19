package main

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	var ret, score int
	for i, j := 0, len(tokens)-1; i <= j && tokens[i] <= power; {
		for i <= j && tokens[i] <= power {
			power -= tokens[i]
			i++
			score += 1
		}
		ret = max(ret, score)
		if score >= 1 {
			score -= 1
			power += tokens[j]
			j--
		}
	}
	return ret
}
