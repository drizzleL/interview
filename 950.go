package main

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	q := make([]int, len(deck))
	for i := range q {
		q[i] = i
	}
	ret := make([]int, len(deck))
	var i int
	for len(q) != 0 {
		ret[q[0]] = deck[i]
		i++
		q = q[1:]
		if len(q) == 0 {
			break
		}
		q = append(q[1:], q[0])
	}
	return ret
}
