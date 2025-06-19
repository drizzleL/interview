package main

func firstUniqChar(s string) int {
	var dict [26]int
	for _, c := range s {
		dict[c-'a'] += 1
	}
	for i, c := range s {
		if dict[c-'a'] != 1 {
			continue
		}
		return i
	}
	return -1
}
