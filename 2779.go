package main

import "sort"

func maximumBeauty(nums []int, k int) int {
	sort.Ints(nums)
	var ret int
	for i := 0; i < len(nums); i++ {
		idx := sort.Search(len(nums)-i, func(j int) bool {
			return nums[i+j] > nums[i]+k*2
		})
		ret = max(ret, idx)
	}
	return ret
}
