package main

import "sort"

func maximizeExpressionOfThree(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-1] + nums[len(nums)-2] - nums[0]
}
