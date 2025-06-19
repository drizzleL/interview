package main

func minGroupsForValidAssignment(balls []int) int {
	dict := map[int]int{}
	for _, b := range balls {
		dict[b] += 1
	}
	minBox := len(balls)
	for _, c := range dict {
		minBox = min(minBox, c)
	}
	check := func(v int) int {
		var ret int
		for _, c := range dict {
			remaining := c % (v + 1)
			g := c / (v + 1)
			if remaining == 0 {
				ret += g
				continue
			}
			if g >= v-remaining {
				ret += g + 1
				continue
			}
			return 0
		}
		return ret
	}
	for i := minBox; i >= 2; i-- {
		v := check(i)
		if v == 0 {
			continue
		}
		return v
	}
	return len(balls)
}
