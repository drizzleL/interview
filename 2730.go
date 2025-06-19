package main

func longestSemiRepetitiveSubstring(s string) int {
	ret := 1
	for l, r, idx := 0, 1, -1; r < len(s); r++ {
		if s[r] == s[r-1] {
			if idx != -1 {
				l = r
			}
			idx = r
		}
		ret = max(ret, r-l+1)
	}
	return ret
}
