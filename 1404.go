package main

func numSteps(s string) int {
	var ret int
	var carry int
	for i := len(s) - 1; i > 0; i-- {
		carry += int(s[i] - '0')
		switch carry {
		case 0:
			ret += 1
		case 1:
			ret += 2
		case 2:
			carry = 1
			ret += 1
		}
	}
	carry += int(s[0] - '0')
	if carry == 2 {
		ret += 1
	}
	return ret
}
