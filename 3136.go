package main

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	var vowel, consonant int
	for _, c := range word {
		if c >= '0' && c <= '9' {
			continue
		}
		if !((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')) {
			return false
		}
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowel += 1
		default:
			consonant += 1
		}
	}
	return vowel > 0 && consonant > 0
}
