package main

func longestSubsequenceRepeatedK(s string, k int) string {
	var cnt [26]int
	for _, c := range s {
		cnt[c-'a'] += 1
	}
	check := func(b []byte) bool {
		var cnt2 [26]int
		for _, c := range b {
			cnt2[c-'a'] += 1
		}
		for i := range cnt {
			if cnt2[i]*k > cnt[i] {
				return false
			}
		}
		var k2 int
		for i, j := 0, 0; i < len(s) && k2 < k; i++ {
			if s[i] != b[j] {
				continue
			}
			j += 1
			if j == len(b) {
				j = 0
				k2 += 1
			}
		}
		return k2 >= k
	}
	var ret string
	cmp := func(b []byte) bool {
		if len(b) < len(ret) {
			return false
		}
		if len(b) > len(ret) {
			return true
		}
		for i := 0; i < len(b); i++ {
			if b[i] == ret[i] {
				continue
			}
			return b[i] > ret[i]
		}
		return true
	}
	var helper func(b []byte)
	helper = func(b []byte) {
		if cmp(b) {
			ret = string(b)
		}
		for j := 0; j < 26; j++ {
			b2 := append(b, byte('a'+j))
			if !check(b2) {
				continue
			}
			helper(b2)
		}
	}
	helper(nil)
	return string(ret)
}
