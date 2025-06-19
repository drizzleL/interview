package main

func hasSpecialSubstring(s string, k int) bool {
	var cnt int
	helper := func(i, j int) bool {
		if i < 0 || j < 0 || i >= len(s) || j >= len(s) {
			return true
		}
		return s[i] != s[j]
	}
	for i := 0; i < len(s); i++ {
		if i != 0 && s[i] == s[i-1] {
			cnt += 1
		} else {
			cnt = 1
		}
		if cnt >= k && helper(i-k, i) && helper(i, i+1) {
			return true
		}
	}
	return false
}
