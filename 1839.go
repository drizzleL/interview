package main

func longestBeautifulSubstring(word string) int {
	dict := map[byte]int{
		'a': 1,
		'e': 2,
		'i': 3,
		'o': 4,
		'u': 5,
	}
	check := func(j int) bool {
		if dict[word[j]] == 0 {
			return false
		}
		return dict[word[j]] == dict[word[j-1]] || dict[word[j]] == dict[word[j-1]]+1
	}
	var ret int
	for j := 0; j < len(word); {
		for j < len(word) && word[j] != 'a' {
			j++
		}
		if j == len(word) {
			break
		}
		i := j
		j += 1
		for j < len(word) && check(j) {
			j++
		}
		if word[j-1] == 'u' {
			ret = max(ret, j-i)
		}
	}
	return ret
}
