package main

import "sort"

func findWinners(matches [][]int) [][]int {
	players := map[int]bool{}
	lost := map[int]int{}
	for _, m := range matches {
		players[m[0]] = true
		players[m[1]] = true
		lost[m[1]] += 1
	}
	var sortedPlayers []int
	for p := range players {
		sortedPlayers = append(sortedPlayers, p)
	}
	sort.Ints(sortedPlayers)
	ret := make([][]int, 2)
	for _, p := range sortedPlayers {
		switch lost[p] {
		case 0:
			ret[0] = append(ret[0], p)
		case 1:
			ret[1] = append(ret[1], p)
		}
	}
	return ret
}
