package main

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	ret := make([]int, len(spells))
	for i, sp := range spells {
		p := int(success) / sp
		if p*sp != int(success) {
			p += 1
		}
		idx := sort.SearchInts(potions, p)
		ret[i] = len(potions) - idx
	}
	return ret
}
