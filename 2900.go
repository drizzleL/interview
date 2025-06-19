package main

func getLongestSubsequence(words []string, groups []int) []string {
	ret := []string{words[0]}
	last := groups[0]
	for i := 1; i < len(words); i++ {
		if groups[i] == last {
			continue
		}
		ret = append(ret, words[i])
		last = groups[i]
	}
	return ret
}
