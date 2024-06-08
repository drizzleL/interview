package main

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for ; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			break
		}
	}
	if i >= j {
		return true
	}
	check := func(s string) bool {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}
	if check(s[i+1:j+1]) || check(s[i:j]) {
		return true
	}
	return false
}
