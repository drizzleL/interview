package main

import (
	"sort"
)

func countWays(nums []int) int {
	sort.Ints(nums)
	var ret int
	if len(nums) != 0 && nums[0] != 0 {
		ret += 1
	}
	for i := 0; i < len(nums); {
		v := nums[i]
		for ; i < len(nums) && nums[i] == v; i++ {
		}
		if v < i && (i == len(nums) || nums[i] > i) {
			ret += 1
		}
	}
	return ret
}
