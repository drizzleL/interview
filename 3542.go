package main

func minOperations16(nums []int) int {
	stack := []int{0}
	var ret int
	for _, num := range nums {
		for stack[len(stack)-1] > num {
			stack = stack[:len(stack)-1]
		}
		if stack[len(stack)-1] < num {
			ret += 1
			stack = append(stack, num)
		}
	}
	return ret
}
