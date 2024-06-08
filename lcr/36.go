package main

import "strconv"

func evalRPN(tokens []string) int {
	vals := []int{}
	for _, token := range tokens {
		switch token {
		case "+":
			a, b := vals[len(vals)-2], vals[len(vals)-1]
			vals = vals[:len(vals)-2]
			vals = append(vals, a+b)
		case "-":
			a, b := vals[len(vals)-2], vals[len(vals)-1]
			vals = vals[:len(vals)-2]
			vals = append(vals, a-b)
		case "*":
			a, b := vals[len(vals)-2], vals[len(vals)-1]
			vals = vals[:len(vals)-2]
			vals = append(vals, a*b)
		case "/":
			a, b := vals[len(vals)-2], vals[len(vals)-1]
			vals = vals[:len(vals)-2]
			vals = append(vals, a/b)
		default:
			d, _ := strconv.Atoi(token)
			vals = append(vals, d)
		}
	}
	return vals[0]
}
