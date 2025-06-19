package main

func minimumLength2(s string) int {
	dict := make([]int, 26)
	for _, c := range dict {
		dict[c] += 1
	}
	var ret int
	for _, v := range dict {
		if v == 0 {
			continue
		}
		ret += 2 - v%2
	}
	return ret
}
