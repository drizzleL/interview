package main

import "sort"

func makeSimilar(nums []int, target []int) int64 {
	var numGroup, targetGroup [2][]int
	for i := range nums {
		key := nums[i] % 2
		numGroup[key] = append(numGroup[key], nums[i])
		key2 := target[i] % 2
		targetGroup[key2] = append(targetGroup[key2], target[i])
	}
	helper := func(a, b []int) int {
		sort.Ints(a)
		sort.Ints(b)
		var ret int
		for i := range a {
			if a[i] > b[i] {
				ret += (a[i] - b[i]) / 2
			}
		}
		return ret
	}
	var ret int
	for i := range numGroup {
		ret += helper(numGroup[i], targetGroup[i])
	}
	return int64(ret)
}
