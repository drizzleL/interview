package main

func resultingString(s string) string {
	var b []byte
	match := func(x byte) bool {
		if len(b) == 0 {
			return false
		}
		c := b[len(b)-1]
		if x > c {
			x, c = c, x
		}
		return c-x == 1 || c-x == 25
	}
	for _, c := range s {
		if !match(byte(c)) {
			b = append(b, byte(c))
			continue
		}
		b = b[:len(b)-1]
	}
	return string(b)
}
