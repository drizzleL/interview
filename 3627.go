package main

import "sort"

func maximumMedianSum(nums []int) int64 {
	sort.Ints(nums)
	var ret int
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-2 {
		ret += nums[j-1]
	}
	return int64(ret)
}
