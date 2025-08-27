package main

import (
	"sort"
)

func minCost8(basket1 []int, basket2 []int) int64 {
	dict := map[int]int{}
	for _, v := range basket1 {
		dict[v] += 1
	}
	for _, v := range basket2 {
		dict[v] -= 1
	}
	minVal := -1
	var swaps []int
	for k := range dict {
		if minVal == -1 || k < minVal {
			minVal = k
		}
		if dict[k]%2 != 0 {
			return -1
		}
		for i := 0; i < abs(dict[k])/2; i++ {
			swaps = append(swaps, k)
		}
	}
	var ret int
	sort.Ints(swaps)
	for i := 0; i < len(swaps)/2; i++ {
		ret += min(minVal*2, swaps[i])
	}
	return int64(ret)
}
