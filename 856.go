package main

func scoreOfParentheses(s string) int {
	var l, ret int
	for i, c := range s {
		if c == '(' {
			l += 1
			continue
		}
		l -= 1
		if s[i-1] == '(' {
			ret += 1 << l
		}
	}
	return ret
}
