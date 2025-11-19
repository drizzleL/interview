package main

import "sort"

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	var ret int
	for i := 0; i < len(nums); i++ {
		var k int
		for j := i + 1; j < len(nums); j++ {
			k = max(k, j+1)
			for k < len(nums) && nums[i]+nums[j] > nums[k] {
				k++
			}
			ret += k - j - 1
		}
	}
	return ret
}
