package main

func countSubstrings(s string) int {
	var ret int
	for i := range s {
		// i as mid
		for l, r := i, i; l >= 0 && r < len(s); l, r = l-1, r+1 {
			if s[l] == s[r] {
				ret++
			} else {
				break
			}
		}
		// i + 1 as mid
		for l, r := i, i+1; l >= 0 && r < len(s); l, r = l-1, r+1 {
			if s[l] == s[r] {
				ret++
			} else {
				break
			}
		}
	}
	return ret
}
