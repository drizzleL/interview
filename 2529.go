package main

import "sort"

func maximumCount(nums []int) int {
	zero := sort.SearchInts(nums, 0)
	if zero == len(nums) {
		return len(nums)
	}
	one := sort.SearchInts(nums, 1)
	pos := len(nums) - one
	return max(zero, pos)
}
