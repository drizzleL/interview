package main

func largestRectangleArea(heights []int) int {
	lh, rh := make([]int, len(heights)), make([]int, len(heights))
	var stack []int
	for i := range heights {
		for len(stack) != 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			lh[i] = i - stack[len(stack)-1] - 1
		} else {
			lh[i] = i
		}
		stack = append(stack, i)
	}
	stack = stack[:0]
	for i := len(heights) - 1; i >= 0; i-- {
		for len(stack) != 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			rh[i] = stack[len(stack)-1] - i - 1
		} else {
			rh[i] = len(heights) - 1 - i
		}
		stack = append(stack, i)
	}
	var ret int
	for i := range heights {
		tmp := heights[i] * (lh[i] + rh[i] + 1)
		if tmp > ret {
			ret = tmp
		}
	}
	return ret
}
