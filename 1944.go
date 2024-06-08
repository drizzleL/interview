package main

import (
	"sort"
)

func canSeePersonsCount(heights []int) []int {
	ret := make([]int, len(heights))
	var stack []int
	for i := len(heights) - 1; i >= 0; i-- {
		idx := sort.Search(len(stack), func(j int) bool {
			return stack[len(stack)-1-j] >= heights[i]
		})
		if idx == len(stack) {
			ret[i] = len(stack)
		} else {
			ret[i] = idx + 1
		}
		for len(stack) != 0 && stack[len(stack)-1] <= heights[i] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, heights[i])
	}
	return ret
}
