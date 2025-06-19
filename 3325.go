package main

func numberOfSubstrings(s string, k int) int {
	var ret int
	dict := [26]int{}
	var i int
	for _, c := range s {
		dict[c-'a'] += 1
		for dict[c-'a'] == k {
			dict[s[i]-'a'] -= 1
			i += 1
		}
		ret += i
	}
	return ret
}
