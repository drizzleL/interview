package main

func isUnique(astr string) bool {
	dict := make([]bool, 26)
	for _, c := range astr {
		if dict[c-'a'] {
			return false
		}
		dict[c-'a'] = true
	}
	return true
}
