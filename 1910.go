package main

func removeOccurrences(s string, part string) string {
	var b []byte
	lastMatch := func() bool {
		if len(b) < len(part) {
			return false
		}
		for i, j := len(part)-1, len(b)-1; i >= 0; i, j = i-1, j-1 {
			if b[j] != part[i] {
				return false
			}
		}
		return true
	}
	for _, c := range s {
		b = append(b, byte(c))
		if !lastMatch() {
			continue
		}
		b = b[:len(b)-len(part)]
	}
	return string(b)
}
