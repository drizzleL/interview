package main

func maximumSubsequenceCount(text string, pattern string) int64 {
	dict := [26]int{}
	dict[pattern[0]-'a'] += 1
	var addPre, addSuff int
	for i := 0; i < len(text); i++ {
		if text[i] == pattern[1] {
			addPre += dict[pattern[0]-'a']
		}
		dict[text[i]-'a'] += 1
	}
	dict = [26]int{}
	dict[pattern[1]-'a'] += 1
	for i := len(text) - 1; i >= 0; i-- {
		if text[i] == pattern[0] {
			addSuff += dict[pattern[1]-'a']
		}
		dict[text[i]-'a'] += 1
	}
	return int64(max(addPre, addSuff))
}
