package main

import "sort"

func findOriginalArray(changed []int) []int {
	if len(changed)%2 != 0 {
		return nil
	}
	dict := map[int]int{}
	for _, v := range changed {
		dict[v] += 1
	}
	sort.Ints(changed)
	ret := make([]int, 0, len(changed)/2)
	for _, v := range changed {
		if dict[v] == 0 {
			continue
		}
		if dict[v*2] == 0 {
			return nil
		}
		ret = append(ret, v)
		dict[v] -= 1
		dict[v*2] -= 1
	}
	return ret
}
