package main

import "fmt"

func getHint(secret string, guess string) string {
	dict := make([]int, 10)
	dict2 := make([]int, 10)
	var bulls, cows int
	for i, c := range guess {
		if byte(c) == secret[i] {
			bulls += 1
			continue
		}
		dict[c-'0'] += 1
		dict2[secret[i]-'0'] += 1
	}
	for i := range dict {
		cows += min(dict[i], dict2[i])
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}
