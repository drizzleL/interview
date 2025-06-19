package main

import "math"

func gridGame(grid [][]int) int64 {
	var sum1, sum2 int
	for i := 0; i < len(grid[0]); i++ {
		sum1 += grid[0][i]
		sum2 += grid[1][i]
	}
	leftSum1, rightSum2 := 0, sum2
	ret := math.MaxInt64
	for i := 0; i < len(grid[0]); i++ {
		leftSum1 += grid[0][i]
		bot2 := max(sum1-leftSum1, sum2-rightSum2)
		ret = min(ret, bot2)
		rightSum2 -= grid[1][i]
	}
	return int64(ret)
}
