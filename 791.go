package main

func customSortString(order string, s string) string {
	dict := make([]int, 26)
	for _, c := range s {
		dict[c-'a']++
	}
	ret := make([]byte, 0, len(s))
	for _, c := range order {
		for i := 0; i < dict[c-'a']; i++ {
			ret = append(ret, byte(c))
		}
		dict[c-'a'] = 0
	}
	for k := range dict {
		c := 'a' + k
		for i := 0; i < dict[c-'a']; i++ {
			ret = append(ret, byte(c))
		}
	}
	return string(ret)
}
