package main

import (
	"math"
	"sort"
)

func minAbsDiff(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	ret := make([][]int, m-k+1)
	for i := range ret {
		ret[i] = make([]int, n-k+1)
	}
	find := func(x []int) int {
		ret := math.MaxInt32
		for i, j := 0, 0; j < len(x); {
			for j < len(x) && x[i] == x[j] {
				j++
			}
			if j == len(x) {
				break
			}
			ret = min(ret, x[j]-x[i])
			i = j
		}
		if ret == math.MaxInt32 {
			return 0
		}
		return ret
	}
	for i := 0; i+k <= m; i++ {
		for j := 0; j+k <= n; j++ {
			var vals []int
			for x := i; x < i+k; x++ {
				for y := j; y < j+k; y++ {
					vals = append(vals, grid[x][y])
				}
			}
			sort.Ints(vals)
			ret[i][j] = find(vals)
		}
	}
	return ret
}
