package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	dict1 := make([]int, 26)
	dict2 := make([]int, 26)
	for _, c := range s1 {
		dict1[c-'a']++
	}
	check := func() bool {
		for i := range dict1 {
			if dict1[i] != dict2[i] {
				return false
			}
		}
		return true
	}
	for i := 0; i < len(s1); i++ {
		c := s2[i]
		dict2[c-'a']++
	}
	if check() {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		c := s2[i]
		oldC := s2[i-len(s1)]
		dict2[c-'a']++
		dict2[oldC-'a']--
		if check() {
			return true
		}
	}
	return false
}
