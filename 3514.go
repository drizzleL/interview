package main

import "math/bits"

func uniqueXorTriplets2(nums []int) int {
	var maxVal int
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	bitLen := bits.Len(uint(maxVal))
	seen := make([]bool, 1<<bitLen)
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			seen[nums[i]^nums[j]] = true
		}
	}
	seen2 := make([]bool, 1<<bitLen)
	var ret int
	for _, num := range nums {
		for k, v := range seen {
			if !v {
				continue
			}
			k2 := num ^ k
			if seen2[k2] {
				continue
			}
			seen2[k2] = true
			ret += 1
		}
	}
	return ret
}
