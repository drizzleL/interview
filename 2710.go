package main

func removeTrailingZeros(num string) string {
	b := []byte(num)
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != '0' {
			break
		}
		b = b[:i]
	}
	return string(b)
}
