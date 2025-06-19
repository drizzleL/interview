package main

func removeStars(s string) string {
	var b []byte
	for _, c := range s {
		if c != '*' {
			b = append(b, byte(c))
			continue
		}
		if len(b) > 0 {
			b = b[:len(b)-1]
		}
	}
	return string(b)
}
