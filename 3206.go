package main

func numberOfAlternatingGroups(colors []int) int {
	var ret int
	for i := 0; i < len(colors); i++ {
		before, after := i-1, i+1
		if i == 0 {
			before = len(colors) - 1
		}
		if i == len(colors)-1 {
			after = 0
		}
		if colors[i] != colors[before] && colors[i] != colors[after] {
			ret += 1
		}
	}
	return ret
}
