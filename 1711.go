package main

import "sort"

func countPairs5(deliciousness []int) int {
	if len(deliciousness) <= 1 {
		return 0
	}
	sort.Ints(deliciousness)
	tmp := deliciousness[len(deliciousness)-1] + deliciousness[len(deliciousness)-2]
	var arr []int
	for v := 1; v <= tmp; v *= 2 {
		arr = append(arr, v)
	}
	dict := map[int]int{}
	var ret int
	for _, dec := range deliciousness {
		for _, v := range arr {
			ret += max(0, dict[v-dec])
		}
		dict[dec] += 1
	}
	return ret
}
