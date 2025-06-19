package main

import "sort"

func maxFrequency3(nums []int, k int) int {
	sort.Ints(nums)
	ret, takes := 1, 0
	for i, left := 1, 0; i < len(nums); i++ {
		takes += (nums[i] - nums[i-1]) * (i - left)
		for takes > k {
			takes -= nums[i] - nums[left]
			left += 1
		}
		ret = max(ret, i-left+1)
	}
	return ret
}
