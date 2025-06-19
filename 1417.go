package main

func reformat(s string) string {
	var b1, b2 []byte
	for _, c := range s {
		if c >= '0' && c <= '9' {
			b1 = append(b1, byte(c))
		} else {
			b2 = append(b2, byte(c))
		}
	}
	if len(b2) > len(b1) {
		b1, b2 = b2, b1
	}
	if len(b1) > len(b2)+1 {
		return ""
	}
	var b []byte
	if len(b1) > len(b2) {
		b = append(b, b1[0])
		b1 = b1[1:]
		b1, b2 = b2, b1
	}
	for i := 0; i < len(b1); i++ {
		b = append(b, b1[i], b2[i])
	}
	return string(b)
}
