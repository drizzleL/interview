package main

import (
	"sort"
)

func incremovableSubarrayCount(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// find increasing tail
	tailStart := len(nums) - 1
	for ; tailStart >= 1; tailStart-- {
		if nums[tailStart] <= nums[tailStart-1] {
			break
		}
	}
	var ret int
	for i := 0; i < len(nums); i++ {
		idx := tailStart
		if i != 0 {
			idx = tailStart + sort.Search(len(nums)-tailStart, func(j int) bool {
				return nums[j+tailStart] > nums[i-1]
			})
		}
		ret += 1 + len(nums) - idx
		if idx == i {
			ret -= 1
		}
		if i >= 1 && nums[i] <= nums[i-1] {
			break
		}
	}
	return ret
}
