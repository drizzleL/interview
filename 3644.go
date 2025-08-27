package main

import "math/bits"

func sortPermutation(nums []int) int {
	maxVal := len(nums) - 1
	mask := 1<<bits.Len(uint(maxVal)) - 1
	if mask != maxVal {
		mask >>= 1
	}
	for i := range nums {
		if nums[i] != i {
			if nums[i] == mask {
				return 0
			}
			return i
		}
	}
	return 0
}
