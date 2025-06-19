package main

import "sort"

func maximizeGreatness(nums []int) int {
	var i int
	sort.Ints(nums)
	for j := 0; j < len(nums); {
		for j < len(nums) && nums[j] == nums[i] { // find first bigger than nums[i]
			j++
		}
		if j == len(nums) {
			return i
		}
		i++
		j++
	}
	return i
}
