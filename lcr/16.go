package main

func lengthOfLongestSubstring(s string) int {
	found := map[byte]bool{}
	var l int
	var ret int
	for i, c := range s {
		for found[byte(c)] {
			found[s[l]] = false
			l++
		}
		found[byte(c)] = true
		if i-l+1 > ret {
			ret = i - l + 1
		}
	}
	return ret
}
