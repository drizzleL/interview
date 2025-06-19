package main

func nextGreaterElements(nums []int) []int {
	var stack []int
	ret := make([]int, len(nums))
	for i, num := range nums {
		ret[i] = num
		for len(stack) > 0 && nums[stack[len(stack)-1]] < num {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret[index] = num
		}
		stack = append(stack, i)
	}
	for i := range ret {
		if ret[i] == nums[i] {
			ret[i] = -1
		}
	}
	for i := 0; len(stack) > 0 && i < len(nums); i++ {
		num := nums[i]
		for len(stack) > 0 && nums[stack[len(stack)-1]] < num {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret[index] = num
		}
	}
	return ret
}
