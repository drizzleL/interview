package main

import "math"

func find132pattern(nums []int) bool {
	var stack [][2]int
	minNum := math.MaxInt32
	for _, num := range nums {
		for len(stack) != 0 && stack[len(stack)-1][1] <= num {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 && stack[len(stack)-1][0] < num {
			return true
		}
		minNum = min(minNum, num)
		stack = append(stack, [2]int{minNum, num})
	}
	return false
}
