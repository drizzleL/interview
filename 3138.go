package main

func minAnagramLength(s string) int {
	check := func(v int) bool {
		var dict [26]int
		for i := 0; i < v; i++ {
			dict[s[i]-'a'] += 1
		}
		for i := v; i < len(s); i += v {
			var dict2 [26]int
			for j := 0; j < v; j++ {
				dict2[s[i+j]-'a'] += 1
			}
			if dict2 != dict {
				return false
			}
		}
		return true
	}
	for v := 1; v < len(s); v++ {
		if len(s)%v != 0 {
			continue
		}
		if !check(v) {
			continue
		}
		return v
	}
	return len(s)
}
