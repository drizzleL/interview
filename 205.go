package main

func isIsomorphic(s string, t string) bool {
	dict := map[byte]byte{}
	dict2 := map[byte]byte{}
	for i := range s {
		dict[s[i]] = t[i]
		dict2[t[i]] = s[i]
	}
	for i := range s {
		if dict[s[i]] != t[i] {
			return false
		}
		if dict2[t[i]] != s[i] {
			return false
		}
	}
	return true
}
