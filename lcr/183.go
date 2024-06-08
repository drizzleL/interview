package main

func maxAltitude(heights []int, limit int) []int {
	if len(heights) == 0 || len(heights) < limit {
		return nil
	}
	stack := []int{}
	var ret []int
	helper := func(i int) {
		for len(stack) != 0 && stack[len(stack)-1] < heights[i] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, heights[i])
	}
	for i := 0; i < limit; i++ {
		helper(i)
	}
	ret = append(ret, stack[0])
	for i := limit; i < len(heights); i++ {
		if len(stack) != 0 && stack[0] == heights[i-limit] {
			stack = stack[1:]
		}
		helper(i)
		ret = append(ret, stack[0])
	}
	return ret
}
