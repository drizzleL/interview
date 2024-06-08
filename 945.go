package main

import "sort"

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	var ret int
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			continue
		}
		ret += nums[i-1] + 1 - nums[i]
		nums[i] = nums[i-1] + 1
	}
	return ret
}
