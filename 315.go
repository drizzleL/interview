package main

import (
	"sort"
)

func countSmaller(nums []int) []int {
	ret := make([]int, len(nums))
	vals := []int{nums[len(nums)-1]}
	for j := len(nums) - 2; j >= 0; j-- {
		idx := sort.SearchInts(vals, nums[j])
		vals = append(vals, 0)
		copy(vals[idx+1:], vals[idx:])
		vals[idx] = nums[j]
		ret[j] = idx
	}
	return ret
}
