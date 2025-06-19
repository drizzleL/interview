package main

import "sort"

func reductionOperations(nums []int) int {
	var ret int
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == nums[0] {
			break
		}
		if nums[i] == nums[i-1] {
			continue
		}
		ret += len(nums) - i
	}
	return ret
}
