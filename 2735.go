package main

import "math"

func minCost2(nums []int, x int) int64 {
	minVal := math.MaxInt64
	mins := make([]int, len(nums))
	for i := range nums {
		mins[i] = nums[i]
		minVal = min(minVal, nums[i])
	}
	copy(mins, nums)
	ret := math.MaxInt64
	for i := 0; i < len(nums); i++ {
		tmp := i * x
		if tmp+minVal*len(nums) >= ret {
			break
		}
		for j := 0; j < len(nums); j++ {
			idx := (j - i + len(nums)) % len(nums)
			mins[j] = min(mins[j], nums[idx])
			tmp += mins[j]
		}
		ret = min(ret, tmp)
	}
	return int64(ret)
}
