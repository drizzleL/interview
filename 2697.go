package main

func makeSmallestPalindrome(s string) string {
	b := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if b[i] == b[j] {
			continue
		}
		if b[i] < b[j] {
			b[j] = b[i]
		} else {
			b[i] = b[j]
		}
	}
	return string(b)
}
