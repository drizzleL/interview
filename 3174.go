package main

func clearDigits(s string) string {
	var b []byte
	for _, c := range s {
		if c >= '0' && c <= '9' {
			if len(b) != 0 && (b[len(b)-1] > '9' || b[len(b)-1] < '0') {
				b = b[:len(b)-1]
			}
			continue
		}
		b = append(b, byte(c))
	}
	return string(b)
}
