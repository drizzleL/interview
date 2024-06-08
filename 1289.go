package main

import (
	"math"
)

func minFallingPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	last := make([]int, len(grid))
	now := make([]int, len(grid))
	leftMax := make([]int, len(grid))
	rightMax := make([]int, len(grid))
	for i := 0; i < len(grid); i++ {
		l, r := math.MaxInt32, math.MaxInt32
		for j := 0; j < len(grid)-1; j++ {
			leftMax[j] = l
			l = min(l, last[j])
			leftMax[j+1] = l
		}
		for j := len(grid) - 1; j > 0; j-- {
			rightMax[j] = r
			r = min(r, last[j])
			rightMax[j-1] = r
		}
		for j := 0; j < len(grid); j++ {
			now[j] = grid[i][j] + min(leftMax[j], rightMax[j])
		}
		now, last = last, now
	}
	ret := math.MaxInt32
	for _, v := range last {
		ret = min(ret, v)
	}
	return ret
}
