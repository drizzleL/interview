package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	scoreCpy := make([]int, len(score))
	copy(scoreCpy, score)
	sort.Ints(scoreCpy)
	dict := map[int]int{}
	for i, c := range scoreCpy {
		dict[c] = len(scoreCpy) - i
	}
	var ret []string
	for _, s := range score {
		n := dict[s]
		var v string
		switch n {
		case 1:
			v = "Gold Medal"
		case 2:
			v = "Silver Medal"
		case 3:
			v = "Bronze Medal"
		default:
			v = strconv.Itoa(n)
		}
		ret = append(ret, v)
	}
	return ret
}
