package main

func minValidStrings(words []string, target string) int {
	helper := func(s string) []int {
		ret := make([]int, len(s))
		for i := 1; i < len(s); i++ {
			j := ret[i-1]
			for j > 0 && s[i] != s[j] {
				j = ret[j-1]
			}
			if s[i] == s[j] {
				j++
			}
			ret[i] = j
		}
		return ret
	}
	p := make([]int, len(target)+1)
	for _, w := range words {
		s := w + "#" + target
		ret := helper(s)
		for i := 1; i <= len(target); i++ {
			p[i] = max(p[i], ret[len(w)+i])
		}
	}
	size := len(target)
	var ret int
	for ; size > 0 && p[size] > 0; ret += 1 {
		size -= p[size]
	}
	if size != 0 {
		return -1
	}
	return ret
}
