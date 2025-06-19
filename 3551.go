package main

import (
	"sort"
)

func minSwaps5(nums []int) int {
	nums2 := make([]int, len(nums))
	copy(nums2, nums)
	digitSum := func(x int) int {
		var ret int
		for x != 0 {
			ret += x % 10
			x /= 10
		}
		return ret
	}
	sort.Slice(nums2, func(i, j int) bool {
		s1, s2 := digitSum(nums2[i]), digitSum(nums2[j])
		if s1 == s2 {
			return nums2[i] < nums2[j]
		}
		return s1 < s2
	})
	dict := map[int]int{}
	for i, v := range nums2 {
		dict[v] = i
	}
	var ret int
	for i := range nums {
		for dict[nums[i]] != i {
			ret += 1
			idx := dict[nums[i]]
			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}
	return ret
}
