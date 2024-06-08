package main

import "sort"

func frequencySort(nums []int) []int {
	dict := map[int]int{}
	for _, num := range nums {
		dict[num] += 1
	}
	sort.Slice(nums, func(i, j int) bool {
		if dict[nums[i]] == dict[nums[j]] {
			return nums[i] > nums[j]
		}
		return dict[nums[i]] < dict[nums[j]]
	})
	return nums
}
