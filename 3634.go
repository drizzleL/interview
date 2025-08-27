package main

import "sort"

func minRemoval(nums []int, k int) int {
	sort.Ints(nums)
	var maxSize int
	for i := 0; i < len(nums); i++ {
		idx := sort.SearchInts(nums, nums[i]*k+1)
		maxSize = max(maxSize, idx-i)
	}
	return len(nums) - maxSize
}
