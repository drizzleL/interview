package main

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	for i, j := 0, 0; i < len(players); i++ {
		for ; j < len(trainers) && trainers[j] < players[i]; j++ {
		}
		if j == len(trainers) {
			return i
		}
		j += 1
	}
	return len(players)
}
