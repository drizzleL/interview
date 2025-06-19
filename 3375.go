package main

import "sort"

func minOperations4(nums []int, k int) int {
	sort.Ints(nums)
	if nums[0] < k {
		return -1
	}
	var ret int
	if nums[0] == k {
		ret = -1
	}
	for i := len(nums) - 1; i >= 0; {
		j := i - 1
		for ; j >= 0 && nums[j] == nums[i]; j-- {
		}
		i = j
		ret += 1
	}
	return ret
}
