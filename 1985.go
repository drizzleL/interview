package main

import "sort"

func kthLargestNumber(nums []string, k int) string {
	sort.Slice(nums, func(i, j int) bool {
		if len(nums[i]) == len(nums[j]) {
			return nums[i] < nums[j]
		}
		return len(nums[i]) < len(nums[j])
	})
	return nums[len(nums)-1-k]
}
