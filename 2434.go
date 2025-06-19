package main

func robotWithString(s string) string {
	var b []byte
	var t []byte
	after := make([]byte, len(s))
	after[len(s)-1] = s[len(s)-1]
	for i := len(s) - 2; i >= 0; i-- {
		after[i] = s[i]
		if after[i] > after[i+1] {
			after[i] = after[i+1]
		}
	}
	for i := 0; i < len(s); i++ {
		for len(t) != 0 && t[len(t)-1] <= after[i] {
			b = append(b, t[len(t)-1])
			t = t[:len(t)-1]
		}
		if s[i] > after[i] {
			t = append(t, s[i])
		} else {
			b = append(b, s[i])
		}
	}
	for i := len(t) - 1; i >= 0; i-- {
		b = append(b, t[i])
	}
	return string(b)
}
