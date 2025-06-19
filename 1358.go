package main

func numberOfSubstrings2(s string) int {
	next := make([][3]int, len(s))
	last := [3]int{len(s), len(s), len(s)}
	for i := len(s) - 1; i >= 0; i-- {
		next[i] = last
		last[s[i]-'a'] = i
	}
	var ret int
	for l, c := range s {
		var r int
		for i, v := range next[l] {
			if i == int(c-'a') {
				continue
			}
			r = max(r, v)
		}
		ret += len(s) - r
	}
	return ret
}
