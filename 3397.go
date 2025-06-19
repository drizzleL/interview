package main

import "sort"

func maxDistinctElements(nums []int, k int) int {
	sort.Ints(nums)
	var ret int
	for i := range nums {
		if i == 0 {
			nums[i] -= k
			ret += 1
			continue
		}
		if nums[i]-k > nums[i-1] {
			nums[i] -= k
			ret += 1
			continue
		}
		if nums[i]+k == nums[i-1] {
			nums[i] += k
			continue
		}
		nums[i] = nums[i-1] + 1
		ret += 1
	}
	return ret
}
