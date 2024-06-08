package main

func isAlienSorted(words []string, order string) bool {
	dict := map[byte]int{}
	for i, c := range order {
		dict[byte(c)] = i
	}
	check := func(w1, w2 string) bool {
		var i int
		for ; i < len(w1) && i < len(w2); i++ {
			if dict[w1[i]] == dict[w2[i]] {
				continue
			}
			return dict[w1[i]] < dict[w2[i]]
		}
		return i == len(w1)
	}
	for i := 0; i < len(words)-1; i++ {
		if !check(words[i], words[i+1]) {
			return false
		}
	}
	return true
}
