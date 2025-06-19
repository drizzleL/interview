package main

import (
	"math/bits"
)

func smallestSubarrays(nums []int) []int {
	var maxVal int
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	bitSize := bits.Len(uint(maxVal))
	dict := make([]int, bitSize)
	ret := make([]int, len(nums))
	j := len(nums) - 1
	check := func(x int) bool {
		for k := 0; k < bitSize; k++ {
			if x&(1<<k) == 0 {
				continue
			}
			if dict[k] <= 1 {
				return false
			}
		}
		return true
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for k := 0; k < bitSize; k++ {
			if nums[i]&(1<<k) != 0 {
				dict[k] += 1
			}
		}
		for j > i && check(nums[j]) {
			for k := 0; k < bitSize; k++ {
				if nums[j]&(1<<k) != 0 {
					dict[k] -= 1
				}
			}
			j -= 1
		}
		ret[i] = j - i + 1
	}
	return ret
}
