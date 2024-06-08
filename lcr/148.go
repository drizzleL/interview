package main

func validateBookSequences(putIn []int, takeOut []int) bool {
	var stack []int
	var i, j int
	for j < len(takeOut) {
		if len(stack) != 0 && stack[len(stack)-1] == takeOut[j] && j < len(takeOut) {
			stack = stack[:len(stack)-1]
			j++
			continue
		}
		if i < len(putIn) {
			stack = append(stack, putIn[i])
			i++
			continue
		}
		return false
	}
	return true
}
