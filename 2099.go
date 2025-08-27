package main

import "sort"

func maxSubsequence(nums []int, k int) []int {
	nums2 := make([]int, len(nums))
	copy(nums2, nums)
	sort.Ints(nums2)
	dict := map[int]int{}
	for i := len(nums2) - 1; i >= len(nums2)-k; i-- {
		dict[nums2[i]] += 1
	}
	var ret []int
	for i := 0; i < len(nums); i++ {
		if dict[nums[i]] > 0 {
			ret = append(ret, nums[i])
			dict[nums[i]] -= 1
		}
	}
	return ret
}
