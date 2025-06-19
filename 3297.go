package main

func validSubstringCount(word1 string, word2 string) int64 {
	var dict, dict2 [26]int
	for _, c := range word2 {
		dict[c-'a'] += 1
	}
	check := func() bool {
		for i := range dict {
			if dict2[i] < dict[i] {
				return false
			}
		}
		return true
	}
	var ret int
	for i, j := 0, 0; i < len(word1); i++ {
		for ; j < len(word1) && !check(); j++ {
			dict2[word1[j]-'a'] += 1
		}
		if check() {
			ret += len(word1) - j
		}
		dict2[word1[i]-'a'] -= 1
	}
	return int64(ret)
}
