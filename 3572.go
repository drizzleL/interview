package main

import "sort"

func maxSumDistinctTriplet(x []int, y []int) int {
	dict := map[int]int{}
	for i, v := range x {
		dict[v] = max(dict[v], y[i])
	}
	var vals []int
	for _, v := range dict {
		vals = append(vals, v)
	}
	sort.Ints(vals)
	if len(vals) < 3 {
		return -1
	}
	return vals[len(vals)-1] + vals[len(vals)-2] + vals[len(vals)-3]
}
