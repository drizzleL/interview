package main

func minimumLength(s string) int {
	l, r := 0, len(s)-1
	for l < r && s[l] == s[r] {
		c := s[l]
		// push l till s[l] not c
		for l <= r && s[l] == c {
			l++
		}
		for l < r && s[r] == c && l < r {
			r--
		}
	}
	return r - l + 1
}
