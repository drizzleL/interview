package main

func smallestPalindrome(s string) string {
	var dict [26]int
	for _, c := range s {
		dict[c-'a'] += 1
	}
	b := make([]byte, len(s))
	for k, v := range dict {
		if v%2 == 1 {
			b[len(b)/2] = byte('a' + k)
			dict[k] -= 1
			break
		}
	}
	for i, j, idx := 0, len(b)-1, 0; i < j; i, j = i+1, j-1 {
		for dict[idx] == 0 {
			idx += 1
		}
		b[i] = byte('a' + idx)
		b[j] = byte('a' + idx)
		dict[idx] -= 2
	}
	return string(b)
}
