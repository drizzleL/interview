package main

func shiftingLetters2(s string, shifts []int) string {
	var cnt int
	b := []byte(s)
	for i := len(s) - 1; i >= 0; i-- {
		d := int(s[i] - 'a')
		cnt += shifts[i]
		cnt %= 26
		d += cnt
		d %= 26
		b[i] = 'a' + byte(d)
	}
	return string(b)
}
