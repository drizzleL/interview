package main

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int
	getNums := func() (int, int) {
		a, b := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		return a, b
	}
	for _, tk := range tokens {
		var num int
		switch tk {
		case "+":
			a, b := getNums()
			num = a + b
		case "-":
			a, b := getNums()
			num = a - b
		case "*":
			a, b := getNums()
			num = a * b
		case "/":
			a, b := getNums()
			num = a / b
		default:
			num, _ = strconv.Atoi(tk)
		}
		stack = append(stack, num)
	}
	return stack[0]
}
