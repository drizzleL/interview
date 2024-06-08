package main

import "sort"

func heightChecker(heights []int) int {
	h2 := make([]int, len(heights))
	copy(h2, heights)
	sort.Ints(h2)
	var ret int
	for i, h := range heights {
		if h != h2[i] {
			ret += 1
		}
	}
	return ret
}
