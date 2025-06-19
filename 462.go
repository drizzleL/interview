package main

import "sort"

func minMoves2(nums []int) int {
	sort.Ints(nums)
	mid := nums[len(nums)/2]
	var ret int
	for _, num := range nums {
		ret += abs(num - mid)
	}
	return ret
}
