package main

func isAlienSorted(words []string, order string) bool {
	dict := make([]int, 26)
	for i, c := range order {
		dict[c-'a'] = i
	}
	cmp := func(a, b string) bool {
		var i, j int
		for ; i < len(a) && j < len(b); i, j = i+1, j+1 {
			if dict[a[i]-'a'] == dict[b[j]-'a'] {
				continue
			}
			return dict[a[i]-'a'] < dict[b[j]-'a']
		}
		return i == len(a)
	}
	for i := 0; i < len(words)-1; i++ {
		if !cmp(words[i], words[i+1]) {
			return false
		}
	}
	return true
}
