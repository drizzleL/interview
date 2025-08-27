package main

func divideString(s string, k int, fill byte) []string {
	var ret []string
	for i := 0; i < len(s); i += k {
		var b []byte
		for j := 0; j < k; j++ {
			if i+j >= len(s) {
				b = append(b, fill)
			} else {
				b = append(b, s[i+j])
			}
		}
		ret = append(ret, string(b))
	}
	return ret
}
