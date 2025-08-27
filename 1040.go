package main

import "sort"

func numMovesStonesII(stones []int) []int {
	sort.Ints(stones)
	ret := make([]int, 2)
	for i := 1; i < len(stones); i++ {
		ret[1] = max(ret[1], stones[i]-stones[i-1]-1)
	}
	return ret
}
