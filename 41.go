package main

import (
	"math"
)

func firstMissingPositive(nums []int) int {
	minVal, maxVal := math.MaxInt32, math.MinInt32
	for i, num := range nums {
		if num <= 0 {
			nums[i] = 0
			continue
		}
		minVal = min(minVal, num)
		maxVal = max(maxVal, num)
	}
	if minVal != 1 {
		return 1
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == minVal+i {
			continue
		}
		if nums[i] == 0 {
			continue
		}
		v := abs(nums[i])
		idx := v - minVal
		if idx < len(nums) {
			if nums[idx] == 0 || nums[idx] == v {
				nums[idx] = v
			}
			if nums[idx] > 0 {
				nums[idx] *= -1
			}
		}
		if nums[i] < 0 {
			nums[i] = minVal + 1
		} else {
			nums[i] = 0
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 { // missing
			return minVal + i
		}
	}
	return maxVal + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
