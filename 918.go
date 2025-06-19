package main

import "math"

func maxSubarraySumCircular(nums []int) int {
	var sum, curr1, curr2, maxSum, minSum int
	maxNum := math.MinInt32
	var foundNonNeg bool
	for _, num := range nums {
		if num >= 0 {
			foundNonNeg = true
		}
		maxNum = max(maxNum, num)
		sum += num
		curr1 += num
		curr1 = max(0, curr1)
		maxSum = max(maxSum, curr1)
		curr2 += num
		curr2 = min(0, curr2)
		minSum = min(minSum, curr2)
	}
	if !foundNonNeg {
		return maxNum
	}
	return max(maxSum, sum-minSum)
}
