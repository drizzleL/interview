package main

func isScramble(s1 string, s2 string) bool {
	cache := map[[2]string]bool{}
	var helper func(s1, s2 string) bool
	helper = func(s1, s2 string) (ret bool) {
		c, ok := cache[[2]string{s1, s2}]
		if ok {
			return c
		}
		defer func() {
			cache[[2]string{s1, s2}] = ret
		}()
		if s1 == s2 {
			return true
		}
		dict := [26]int{}
		for _, c := range s1 {
			dict[c-'a']++
		}
		for _, c := range s2 {
			dict[c-'a']--
			if dict[c-'a'] < 0 {
				return false
			}
		}
		for i := 1; i < len(s1); i++ {
			a, b := s1[:i], s1[i:]
			a2, b2 := s2[:i], s2[i:]
			a3, b3 := s2[len(s1)-i:], s2[:len(s1)-i]
			if (helper(a, a2) && helper(b, b2)) || (helper(a, a3) && helper(b, b3)) {
				return true
			}
		}
		return false
	}
	return helper(s1, s2)
}
