package main

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	var ret int
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		ret = max(ret, nums[i]+nums[j])
	}
	return ret
}
