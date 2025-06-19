package main

import (
	"sort"
)

func minMaxSums(nums []int, k int) int {
	sort.Ints(nums)
	var ret int
	s := 1
	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
		ret += (nums[i] + nums[j]) * s
		ret %= 1e9 + 7
		s = s*2 - combination(i, k-1)
	}
	return ret
}
