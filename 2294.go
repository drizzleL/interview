package main

import "sort"

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	leftIdx := 0
	ret := 1
	for i, num := range nums {
		if num-nums[leftIdx] <= k {
			continue
		}
		leftIdx = i
		ret += 1
	}
	return ret
}
