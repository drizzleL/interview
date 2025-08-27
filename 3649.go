package main

import "sort"

func perfectPairs(nums []int) int64 {
	for i, num := range nums {
		nums[i] = abs(num)
	}
	sort.Ints(nums)
	var ret int
	for i, j := 0, 0; i < len(nums); i++ {
		for j < len(nums) && nums[j] <= nums[i]*2 {
			j++
		}
		ret += j - i - 1
	}
	return int64(ret)
}
