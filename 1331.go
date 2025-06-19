package main

import "sort"

func arrayRankTransform(arr []int) []int {
	cpy := make([]int, len(arr))
	copy(cpy, arr)
	sort.Ints(cpy)
	dict := map[int]int{}
	for i, rank := 0, 0; i < len(cpy); i++ {
		if i != 0 && cpy[i] == cpy[i-1] {
			continue
		}
		rank += 1
		dict[cpy[i]] = rank
	}
	ret := make([]int, len(arr))
	for i, num := range arr {
		ret[i] = dict[num]
	}
	return ret
}
