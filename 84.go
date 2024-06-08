package main

func largestRectangleArea(heights []int) int {
	var ret int
	heights = append(heights, 0)
	stack := []int{}
	for i, v := range heights {
		for len(stack) != 0 && v < heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h := heights[top]
			if len(stack) == 0 {
				ret = max(ret, h*i)
			} else {
				ret = max(ret, h*(i-1-stack[len(stack)-1]))
			}
		}
		if len(stack) == 0 || v >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}
	return ret
}
