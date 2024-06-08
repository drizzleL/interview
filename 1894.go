package main

import (
	"sort"
)

func chalkReplacer(chalk []int, k int) int {
	var sum int
	for i, c := range chalk {
		sum += c
		if i != 0 {
			chalk[i] += chalk[i-1]
		}
	}
	k %= sum
	return sort.SearchInts(chalk, k+1)
}
