package main

func makeFancyString(s string) string {
	var b []byte
	var cnt int
	for i := 0; i < len(s); i++ {
		if i != 0 && s[i] == s[i-1] {
			if cnt >= 2 {
				continue
			}
		} else {
			cnt = 0
		}
		cnt += 1
		b = append(b, s[i])
	}
	return string(b)
}
