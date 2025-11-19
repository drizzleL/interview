package main

import "sort"

func findMissingElements(nums []int) []int {
	sort.Ints(nums)
	var ret []int
	v := nums[0]
	for _, num := range nums {
		for v != num {
			ret = append(ret, v)
			v += 1
		}
		v += 1
	}
	return ret
}
