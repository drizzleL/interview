package main

import "sort"

func findPairs(nums []int, k int) int {
	sort.Ints(nums)
	dict := map[int]bool{}
	var ret int
	for _, num := range nums {
		ok, found := dict[num-k]
		if !found {
			dict[num] = true
			continue
		}
		if !ok {
			continue
		}
		ret += 1
		dict[num] = true
		dict[num-k] = false
	}
	return ret
}
