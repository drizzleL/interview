package main

import "sort"

func maxAlternatingSum2(nums []int) int64 {
	for i, num := range nums {
		nums[i] *= num
	}
	sort.Ints(nums)
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	var ret int
	half := (len(nums) + 1) / 2
	for i := 0; i < half; i += 1 {
		ret += nums[i]
	}
	for i := half; i < len(nums); i += 1 {
		ret -= nums[i]
	}
	return int64(ret)
}
