package main

import "math"

func maxSumTrionic(nums []int) int64 {
	ret := math.MinInt64
	for i := 1; i+1 < len(nums); i++ { // find decr
		if nums[i+1] >= nums[i] {
			continue
		}
		start := i
		var sum int
		for ; i+1 < len(nums); i++ {
			sum += nums[i]
			if nums[i+1] >= nums[i] {
				break
			}
		}
		// check left
		if nums[start-1] == nums[start] {
			continue
		}
		leftSum, leftSumMax := 0, math.MinInt64
		for j := start - 1; j >= 0 && nums[j] < nums[j+1]; j-- {
			leftSum += nums[j]
			leftSumMax = max(leftSumMax, leftSum)
		}
		// check right
		if i+1 == len(nums) || nums[i+1] == nums[i] {
			continue
		}
		rightSum, rightSumMax := 0, math.MinInt64
		for j := i + 1; j < len(nums) && nums[j] > nums[j-1]; j++ {
			rightSum += nums[j]
			rightSumMax = max(rightSumMax, rightSum)
		}
		sum += leftSumMax + rightSumMax
		ret = max(ret, sum)
	}
	return int64(ret)
}
