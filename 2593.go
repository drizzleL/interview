package main

import (
	"sort"
)

func findScore(nums []int) int64 {
	eles := make([]int, len(nums))
	for i := range nums {
		eles[i] = i
	}
	sort.Slice(eles, func(i, j int) bool {
		if nums[eles[i]] == nums[eles[j]] {
			return eles[i] < eles[j]
		}
		return nums[eles[i]] < nums[eles[j]]
	})
	var ret int
	for i := 0; i < len(eles); i++ {
		ele := eles[i]
		if nums[ele] == 0 {
			continue
		}
		ret += nums[ele]
		nums[ele] = 0
		if ele-1 >= 0 {
			nums[ele-1] = 0
		}
		if ele+1 < len(eles) {
			nums[ele+1] = 0
		}
	}
	return int64(ret)
}
