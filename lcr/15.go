package main

func findAnagrams(s2 string, s1 string) []int {
	if len(s2) < len(s1) {
		return nil
	}
	dict := [26]int{}
	for _, c := range s1 {
		dict[c-'a']++
	}
	check := func() bool {
		for _, v := range dict {
			if v != 0 {
				return false
			}
		}
		return true
	}
	for i := 0; i < len(s1); i++ {
		dict[s2[i]-'a']--
	}
	var ret []int
	if check() {
		ret = append(ret, 0)
	}
	for j := len(s1); j < len(s2); j++ {
		dict[s2[j]-'a']--
		dict[s2[j-len(s1)]-'a']++
		if check() {
			ret = append(ret, j-len(s1)+1)
		}
	}
	return ret
}
