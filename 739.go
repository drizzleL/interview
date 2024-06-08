package main

func dailyTemperatures(temperatures []int) []int {
	ret := make([]int, len(temperatures))
	var stack []int
	for i, temp := range temperatures {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] < temp {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret[top] = i - top
		}
		stack = append(stack, i)
	}
	return ret
}
