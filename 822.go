package main

func flipgame(fronts []int, backs []int) int {
	dict := make([]int, 2001)
	for i := range fronts {
		if fronts[i] == backs[i] {
			dict[fronts[i]] = 2
			continue
		}
		if dict[fronts[i]] != 2 {
			dict[fronts[i]] = 1
		}
		if dict[backs[i]] != 2 {
			dict[backs[i]] = 1
		}
	}
	for i := 1; i < len(dict); i++ {
		if dict[i] != 1 {
			continue
		}
		return i
	}
	return 0
}
