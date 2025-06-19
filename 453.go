package main

import "sort"

func minMoves8(nums []int) int {
	sort.Ints(nums)
	var ret int
	for _, num := range nums {
		ret += num - nums[0]
	}
	return ret
}
