package main

import "sort"

func triangleType(nums []int) string {
	sort.Ints(nums)
	if nums[2] >= nums[0]+nums[1] {
		return "none"
	}
	if nums[0] != nums[1] && nums[1] != nums[2] {
		return "scalene"
	}
	if nums[1] == nums[0] && nums[1] == nums[2] {
		return "equilateral"
	}
	return "isosceles"
}
