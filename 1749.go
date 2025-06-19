package main

import "math"

func maxAbsoluteSum(nums []int) int {
	var leftMax, leftMin int
	ret := math.MinInt32
	var leftSum int
	for _, num := range nums {
		leftSum += num
		ret = max(ret, leftSum-leftMin)
		ret = max(ret, leftMax-leftSum)
		leftMax = max(leftMax, leftSum)
		leftMin = min(leftMin, leftSum)
	}
	return ret
}
