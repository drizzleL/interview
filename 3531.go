package main

func countCoveredBuildings(n int, buildings [][]int) int {
	left := map[int]int{}
	right := map[int]int{}
	up := map[int]int{}
	down := map[int]int{}
	for i, bd := range buildings {
		if old, ok := left[bd[1]]; !ok || bd[0] < buildings[old][0] {
			left[bd[1]] = i
		}
		if old, ok := right[bd[1]]; !ok || bd[0] > buildings[old][0] {
			right[bd[1]] = i
		}
		if old, ok := up[bd[0]]; !ok || bd[1] > buildings[old][1] {
			up[bd[0]] = i
		}
		if old, ok := down[bd[0]]; !ok || bd[1] < buildings[old][1] {
			down[bd[0]] = i
		}
	}
	var ret int
	for i, bd := range buildings {
		if left[bd[1]] == i {
			continue
		}
		if right[bd[1]] == i {
			continue
		}
		if up[bd[0]] == i {
			continue
		}
		if down[bd[0]] == i {
			continue
		}
		ret += 1
	}
	return ret
}
