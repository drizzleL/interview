package main

func numEquivDominoPairs(dominoes [][]int) int {
	dict := map[[2]int]int{}
	var ret int
	for _, d := range dominoes {
		a, b := d[0], d[1]
		if a > b {
			a, b = b, a
		}
		ret += dict[[2]int{a, b}]
		dict[[2]int{a, b}] += 1
	}
	return ret
}
