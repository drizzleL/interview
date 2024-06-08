package main

import "math"

func findDuplicate(nums []int) int {
	minVal, maxVal := math.MaxInt32, math.MinInt32
	for _, num := range nums {
		minVal = min(minVal, num)
		maxVal = min(maxVal, num)
	}
	seen := make([]bool, maxVal-minVal+1)
	for _, num := range nums {
		idx := num - minVal
		if seen[idx] {
			return minVal + idx
		}
		seen[idx] = true
	}
	return -1
}
