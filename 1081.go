package main

func smallestSubsequence(s string) string {
	dict := make([]int, 26)
	for _, c := range s {
		dict[c-'a']++
	}
	var b []byte
	findInb := func(target rune) int {
		for i, c := range b {
			if c == byte(target) {
				return i
			}
		}
		return -1
	}
	needChange := func(idx int, target byte) bool {
		for i := idx + 1; i < len(b); i++ {
			c := b[i]
			if c < target {
				return true
			}
			if dict[c-'a'] == 0 {
				return false
			}
		}
		return false
	}
	for _, c := range s {
		dict[c-'a']--
		idx := findInb(c)
		if idx == -1 {
			b = append(b, byte(c))
			continue
		}
		if needChange(idx, byte(c)) {
			b = append(b[:idx], b[idx+1:]...)
			b = append(b, byte(c))
		}
	}
	return string(b)
}
