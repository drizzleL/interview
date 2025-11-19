package main

func minOperations72(s string) int {
	var dict [26]bool
	for _, c := range s {
		dict[c-'a'] = true
	}
	for i := 1; i < 26; i++ {
		if dict[i] {
			return 26 - i
		}
	}
	return 0
}
