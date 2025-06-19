package main

import "sort"

func lexicographicallySmallestArray(nums []int, limit int) []int {
	cpy := make([]int, len(nums))
	copy(cpy, nums)
	sort.Ints(cpy)
	numGroup := map[int]int{}
	var groups [][]int
	for i, num := range cpy {
		var groupId int
		if i == 0 || cpy[i-1]+limit < num {
			groupId = len(groups)
			groups = append(groups, nil)
		} else {
			groupId = len(groups) - 1
		}
		numGroup[num] = groupId
		groups[groupId] = append(groups[groupId], num)
	}
	ret := make([]int, len(nums))
	for i, num := range nums {
		groupId := numGroup[num]
		ret[i] = groups[groupId][0]
		groups[groupId] = groups[groupId][1:]
	}
	return ret
}
