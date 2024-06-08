package main

func reorganizeString(s string) string {
	dict := make([]int, 26)
	var maxI int
	for _, c := range s {
		idx := int(c - 'a')
		dict[idx]++
		if dict[idx] > dict[maxI] {
			maxI = idx
		}
	}
	if dict[maxI] > (len(s)+1)/2 {
		return ""
	}
	ret := make([]byte, len(s))
	var idx int
	for dict[maxI] != 0 {
		ret[idx] = byte('a' + maxI)
		dict[maxI]--
		idx += 2
		if idx >= len(ret) {
			idx = 1
		}
	}
	for i := range dict {
		for dict[i] > 0 {
			ret[idx] = byte('a' + i)
			dict[i]--
			idx += 2
			if idx >= len(ret) {
				idx = 1
			}
		}
	}
	return string(ret)
}
