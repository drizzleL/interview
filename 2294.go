package main

import "sort"

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	ret := 1
	for last, i := nums[0], 1; i < len(nums); i++ {
		if nums[i]-last > k {
			ret++
			last = nums[i]
		}
	}
	return ret
}
