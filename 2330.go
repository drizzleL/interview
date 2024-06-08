package main

func makePalindrome(s string) bool {
	leftChance := 2
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] == s[j] {
			continue
		}
		if leftChance == 0 {
			return false
		}
		leftChance -= 1
	}
	return true
}
