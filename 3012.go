package main

import "sort"

func minimumArrayLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	var i int
	for i < len(nums) && nums[i] == nums[0] {
		i++
	}
	if i == 1 {
		return 1
	}
	for j := i; j < len(nums); j++ {
		if nums[j]%nums[0] != 0 {
			return 1
		}
	}
	return (i + 1) / 2
}
