package main

func numTilePossibilities(tiles string) int {
	dict := make([]int, 26)
	for _, t := range tiles {
		dict[t-'A'] += 1
	}
	var ret int
	var helper func() int
	helper = func() int {
		var sum int
		for i := range dict {
			if dict[i] == 0 {
				continue
			}
			ret += 1
			dict[i] -= 1
			sum += helper()
			dict[i] += 1
		}
		return sum
	}
	return helper()
}
