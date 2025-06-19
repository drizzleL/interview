package main

import "fmt"

func greatestLetter(s string) string {
	var lower, upper [26]bool
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			lower[c-'a'] = true
		} else {
			upper[c-'A'] = true
		}
	}
	for i := 25; i >= 0; i-- {
		if upper[i] && lower[i] {
			return fmt.Sprintf("%c", 'A'+i)
		}
	}
	return ""
}
