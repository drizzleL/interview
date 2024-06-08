package main

import "math"

func maxValueAfterReverse(nums []int) int {
	presubSums := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		presubSums[i] = presubSums[i-1] + abs(nums[i]-nums[i-1])
	}
	subSum := func(i, j int) int {
		return presubSums[j] - presubSums[i]
	}
	maxi, mini := math.MinInt32, math.MaxInt32
	for i := 1; i < len(nums); i++ {
		a, b := nums[i-1], nums[i]
		maxi = max(maxi, min(a, b))
		mini = min(mini, max(a, b))
	}
	change := max(0, (maxi-mini)*2)
	for i := 1; i < len(nums); i++ {
		a, b := nums[i-1], nums[i]
		tmp1 := -abs(a-b) + abs(nums[0]-b)
		tmp2 := -abs(a-b) + abs(nums[len(nums)-1]-a)
		change = max(change, tmp1)
		change = max(change, tmp2)
	}
	return subSum(0, len(nums)-1) + change
}
