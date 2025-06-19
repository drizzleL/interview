package main

import "sort"

func smallestRangeII(nums []int, k int) int {
	sort.Ints(nums)
	left, right := nums[0]+k, nums[len(nums)-1]-k
	ret := nums[len(nums)-1] - nums[0]
	for i := 0; i < len(nums)-1; i++ {
		minVal := min(left, nums[i+1]-k)
		maxVal := max(right, nums[i]+k)
		ret = min(ret, maxVal-minVal)
	}
	return ret
}
