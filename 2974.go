package main

import "sort"

func numberGame(nums []int) []int {
	sort.Ints(nums)
	for i := 0; i*2 < len(nums); i++ {
		nums[i*2], nums[i*2+1] = nums[i*2+1], nums[i*2]
	}
	return nums
}
