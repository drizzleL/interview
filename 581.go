package main

import (
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	start := 1
	for ; start < len(nums); start++ {
		if nums[start] < nums[start-1] {
			break
		}
	}
	if start == len(nums) { // all asc
		return 0
	}
	end := len(nums) - 2
	for ; ; end-- {
		if nums[end] > nums[end+1] {
			break
		}
	}
	minVal, maxVal := nums[start], nums[end]
	for i := start; i < len(nums); i++ {
		minVal = min(minVal, nums[i])
	}
	for i := end; i >= 0; i-- {
		maxVal = max(maxVal, nums[i])
	}
	startIdx := sort.SearchInts(nums[:start], minVal)
	endIdx := end + sort.SearchInts(nums[end+1:], maxVal) + 1
	return endIdx - startIdx
}
