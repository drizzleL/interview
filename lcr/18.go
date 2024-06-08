package main

func isPalindrome(s string) bool {
	isCh := func(c byte) bool {
		if c >= '0' && c <= '9' {
			return true
		}
		if c >= 'a' && c <= 'z' {
			return true
		}
		if c >= 'A' && c <= 'Z' {
			return true
		}
		return false
	}
	formatCh := func(c byte) byte {
		if c >= 'A' && c <= 'Z' {
			return c - 'A' + 'a'
		}
		return c
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for i < j && !isCh(s[i]) {
			i++
		}
		if i == j {
			return true
		}
		for i < j && !isCh(s[j]) {
			j--
		}
		if i == j {
			return true
		}
		if formatCh(s[i]) != formatCh(s[j]) {
			return false
		}
	}
	return true
}
