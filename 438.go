package main

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}
	var curr, pattern [26]int
	for i := 0; i < len(p); i++ {
		pattern[p[i]-'a'] += 1
	}
	var ret []int
	for i := 0; i < len(p)-1; i++ {
		curr[s[i]-'a'] += 1
	}
	for i := 0; i+len(p) <= len(s); i++ {
		curr[s[i+len(p)-1]-'a'] += 1
		if curr == pattern {
			ret = append(ret, i)
		}
		curr[s[i]-'a'] -= 1
	}
	return ret
}
