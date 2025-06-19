package main

import (
	"sort"
)

func minimalKSum(nums []int, k int) int64 {
	sort.Ints(nums)
	var ret int
	start := 1
	for i := 0; k != 0 && i < len(nums); i++ {
		if nums[i] < start {
			continue
		}
		used := min(k, nums[i]-start)
		end := start + used - 1
		ret += (start + end) * used / 2
		k -= used
		start = nums[i] + 1
	}
	if k != 0 {
		ret += (start*2 + k - 1) * k / 2
	}
	return int64(ret)
}
