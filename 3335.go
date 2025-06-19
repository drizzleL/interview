package main

func lengthAfterTransformations(s string, t int) int {
	var dict [26]int
	for _, c := range s {
		dict[c-'a'] += 1
	}
	var start int
	trans := func() {
		end := start - 1
		if end < 0 {
			end += 26
		}
		dict[start] += dict[end]
		dict[start] %= 1e9 + 7
		start = (start + 25) % 26
	}
	for i := 0; i < t; i++ {
		trans()
	}
	var ret int
	for _, v := range dict {
		ret += v
		ret %= 1e9 + 7
	}
	return ret
}
