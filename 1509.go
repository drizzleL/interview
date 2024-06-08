package main

import (
	"math"
	"sort"
)

func minDifference(nums []int) int {
	if len(nums) <= 3 {
		return 0
	}
	sort.Ints(nums)
	ret := math.MaxInt32
	for l := 0; l <= 3; l++ {
		r := 3 - l
		v1, v2 := nums[l], nums[len(nums)-1-r]
		ret = min(ret, v2-v1)
	}
	return ret
}
