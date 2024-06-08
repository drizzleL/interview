package main

import (
	"math"
	"sort"
)

func findValueOfPartition(nums []int) int {
	sort.Ints(nums)
	ret := math.MaxInt32
	for i := 0; i < len(nums)-1; i++ {
		ret = min(ret, nums[i+1]-nums[i])
	}
	return ret
}
