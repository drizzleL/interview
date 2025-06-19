package main

import "sort"

func numMovesStones(a int, b int, c int) []int {
	vals := []int{a, b, c}
	sort.Ints(vals)
	var cnt int
	if vals[0]+2 == vals[2] {
		cnt = 0
	} else if vals[0]+2 <= vals[1] || vals[1]+2 <= vals[2] {
		cnt = 1
	} else {
		cnt = 2
	}
	return []int{cnt, vals[2] - vals[0] - 2}
}
