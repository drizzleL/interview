package main

func possibleStringCount(word string) int {
	ret := 1
	for i := 0; i < len(word); {
		i += 1
		for i < len(word) && word[i] == word[i-1] {
			ret += 1
		}
	}
	return ret
}
