package main

func findGoodStrings(n int, s1 string, s2 string, evil string) int {
	getCnt := func(s string) int {
		ret := 1
		base := 1
		mod := int(1e9 + 7)
		for i := len(s) - 1; i >= 0; i, base = i-1, (base*26)%mod {
			c := int(s[i] - 'a')
			ret += c * base
			ret %= mod
		}
		return ret
	}
	ret := getCnt(s2) - getCnt(s1) + 1
	for i := 0; i < n; i++ {
		if s1[i] == s2[i] {
			continue
		}
	}
	return ret
}
