package main

import "sort"

func maximumPoints(enemyEnergies []int, currentEnergy int) int64 {
	sort.Ints(enemyEnergies)
	if currentEnergy < enemyEnergies[0] {
		return 0
	}
	var sum int
	for j := 1; j < len(enemyEnergies); j++ {
		sum += enemyEnergies[j]
	}
	return int64((sum + currentEnergy) / enemyEnergies[0])
}
