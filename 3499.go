package main

func maxActiveSectionsAfterTrade(s string) int {
	var ret int
	var extra int
	var l, r int
	for i, c := range s {
		if c == '0' {
			r += 1
			continue
		}
		ret += 1
		if i != 0 && s[i] == '0' {
			if l != 0 && r != 0 {
				extra = max(extra, l+r)
			}
			l = r
			r = 0
		}
	}
	if l != 0 && r != 0 {
		extra = max(extra, l+r)
	}
	return ret + extra
}
