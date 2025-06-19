package main

func similarPairs(words []string) int {
	var ret int
	dict := map[int]int{}
	for _, w := range words {
		var id int
		for _, c := range w {
			id |= 1 << int(c-'a')
		}
		ret += dict[id]
		dict[id] += 1
	}
	return ret
}
