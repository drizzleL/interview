package main

func numberOfSpecialChars(word string) int {
	var upper, lower [26][2]int
	for i := range upper {
		upper[i] = [2]int{-1, -1}
		lower[i] = [2]int{-1, -1}
	}
	for i, c := range word {
		switch {
		case c >= 'a' && c <= 'z':
			if lower[c-'a'][0] == -1 {
				lower[c-'a'][0] = i
			}
			lower[c-'a'][1] = i
		case c >= 'A' && c <= 'Z':
			if upper[c-'A'][0] == -1 {
				upper[c-'A'][0] = i
			}
			upper[c-'A'][1] = i
		}
	}
	var ret int
	for i := range upper {
		if upper[i][0] == -1 || lower[i][0] == -1 {
			continue
		}
		if upper[i][0] > lower[i][1] {
			ret += 1
		}
	}
	return ret
}
