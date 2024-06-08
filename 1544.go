package main

func makeGood(s string) string {
	var stack []rune
	for _, c := range s {
		stack = append(stack, c)
		for len(stack) >= 2 {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			if a-b == 'A'-'a' || a-b == 'a'-'A' {
				stack = stack[:len(stack)-2]
			}
		}
	}
	return string(stack)
}
