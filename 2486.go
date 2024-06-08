package main

func appendCharacters(s string, t string) int {
	var i int
	for j := 0; i < len(t) && j < len(s); j++ {
		if s[j] == t[i] {
			i++
		}
	}
	return len(t) - i
}
