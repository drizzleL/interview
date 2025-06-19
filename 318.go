package main

func maxProduct4(words []string) int {
	var ret int
	ids := make([]int, len(words))
	for i, w := range words {
		var id int
		for _, c := range w {
			id |= 1 << (c - 'a')
		}
		ids[i] = id
	}
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if ids[i]&ids[j] != 0 {
				continue
			}
			ret = max(ret, len(words[i])*len(words[j]))
		}
	}
	return ret
}
