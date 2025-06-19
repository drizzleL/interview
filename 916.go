package main

func wordSubsets(words1 []string, words2 []string) []string {
	helper := func(w string) [26]int {
		var dict [26]int
		for _, c := range w {
			dict[c-'a'] += 1
		}
		return dict
	}
	var maxDict [26]int
	for _, w := range words2 {
		d := helper(w)
		for i := range d {
			maxDict[i] = max(maxDict[i], d[i])
		}
	}
	var ret []string
	for _, w := range words1 {
		d := helper(w)
		var flag bool
		for i := range d {
			if d[i] >= maxDict[i] {
				continue
			}
			flag = true
		}
		if flag {
			continue
		}
		ret = append(ret, w)
	}
	return ret
}
