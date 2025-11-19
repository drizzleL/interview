package main

import "sort"

func maxProduct8(nums []int) int64 {
	for i := range nums {
		nums[i] = abs(nums[i])
	}
	sort.Ints(nums)
	vals := nums[len(nums)-3:]
	if vals[1] == 0 {
		return 0
	}
	return int64(vals[2] * vals[1] * 1e5)
}
