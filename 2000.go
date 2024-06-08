package main

func reversePrefix(word string, ch byte) string {
	idx := -1
	for i := range word {
		if word[i] != ch {
			continue
		}
		idx = i
		break
	}
	if idx == -1 {
		return word
	}
	b := []byte(word)
	for i, j := 0, idx; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
