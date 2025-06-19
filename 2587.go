package main

import "sort"

func maxScore5(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	var sum int
	var ret int
	for _, num := range nums {
		sum += num
		if sum > 0 {
			ret += 1
		} else {
			break
		}
	}
	return ret
}
