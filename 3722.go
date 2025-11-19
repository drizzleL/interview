package main

func lexSmallest(s string) string {
	ret := s
	reverse := func(s string) string {
		b := []byte(s)
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
		return string(b)
	}
	for k := 2; k <= len(s); k++ {
		tmp := reverse(s[:k]) + s[k:]
		if tmp < ret {
			ret = tmp
		}
		tmp = s[:len(s)-k] + reverse(s[len(s)-k:])
		if tmp < ret {
			ret = tmp
		}
	}
	return ret
}
