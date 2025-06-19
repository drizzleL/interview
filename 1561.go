package main

import "sort"

func maxCoins(piles []int) int {
	sort.Ints(piles)
	var ret int
	for i, j := 0, len(piles)-2; i < j; i, j = i+1, j-2 {
		ret += piles[j]
	}
	return ret
}
