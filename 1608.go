package main

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		leftNum := len(nums) - i
		if v >= leftNum && (i == 0 || nums[i-1] < leftNum) {
			return leftNum
		}
	}
	return -1
}
