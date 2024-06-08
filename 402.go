package main

func removeKdigits(num string, k int) string {
	var stack []rune
	for _, c := range num {
		for k > 0 && len(stack) > 0 && c < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k -= 1
		}
		stack = append(stack, c)
	}
	for len(stack) != 0 && stack[0] == '0' {
		stack = stack[1:]
	}
	stack = stack[:len(stack)-k]
	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}
