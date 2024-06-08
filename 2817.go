package main

import (
	"math"
	"sort"
)

func minAbsoluteDifference(nums []int, x int) int {
	if len(nums) == 0 {
		return 0
	}
	sortedNums := []int{nums[0]}
	ret := math.MaxInt32
	for i := x; i < len(nums); i++ {
		nearest := sort.SearchInts(sortedNums, nums[i])
		if nearest < len(sortedNums) {
			ret = min(ret, sortedNums[nearest]-nums[i])
		} else {
			ret = min(ret, nums[i]-sortedNums[len(sortedNums)-1])
		}
		if nearest != 0 {
			ret = min(ret, nums[i]-sortedNums[nearest-1])
		}
		j := i - x + 1
		idx := sort.SearchInts(sortedNums, nums[j])
		if idx < len(sortedNums) && sortedNums[idx] == nums[j] { // existed, skip
			continue
		}
		if idx == len(sortedNums) {
			sortedNums = append(sortedNums, nums[j])
			continue
		}
		sortedNums = append(sortedNums, 0)
		copy(sortedNums[idx+1:], sortedNums[idx:])
		sortedNums[idx] = nums[j]
	}
	return ret
}
