package main

import "sort"

func minOperations14(grid [][]int, x int) int {
	var vals []int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			vals = append(vals, grid[i][j])
		}
	}
	sort.Ints(vals)
	median := vals[len(vals)/2]
	var ret int
	for i := range vals {
		diff := abs(vals[i] - median)
		if diff%x != 0 {
			return -1
		}
		ret += diff / x
	}
	return ret
}
