package main

import "sort"

func maxKDistinct(nums []int, k int) []int {
	sort.Ints(nums)
	var ret []int
	for i := len(nums) - 1; i >= 0 && len(ret) < k; i-- {
		if len(ret) == 0 || ret[len(ret)-1] != nums[i] {
			ret = append(ret, nums[i])
			continue
		}
	}
	return ret
}
