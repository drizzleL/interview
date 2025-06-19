package main

func canBeValid(s string, locked string) bool {
	if len(s)%2 != 0 {
		return false
	}
	var v int
	for i := 0; i < len(s); i++ {
		if locked[i] == '0' {
			v += 1
			continue
		}
		if s[i] == ')' {
			if v == 0 {
				return false
			}
			v -= 1
		} else {
			v += 1
		}
	}
	v = 0
	for i := len(s) - 1; i >= 0; i-- {
		if locked[i] == '0' {
			v += 1
			continue
		}
		if s[i] == '(' {
			if v == 0 {
				return false
			}
			v -= 1
		} else {
			v += 1
		}
	}
	return true
}
