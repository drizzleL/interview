package main

func commonChars(words []string) []string {
	if len(words) == 0 {
		return nil
	}
	helper := func(word string) [26]int {
		var ret [26]int
		for _, c := range word {
			ret[c-'a'] += 1
		}
		return ret
	}
	dict := helper(words[0])
	for i := 1; i < len(words); i++ {
		dict2 := helper(words[i])
		for i := range dict {
			dict[i] = min(dict[i], dict2[i])
		}
	}
	var ret []string
	for i, c := range dict {
		for j := 0; j < c; j++ {
			ret = append(ret, string('a'+i))
		}
	}
	return ret
}
