package main

import "math"

func earliestTime(tasks [][]int) int {
	ret := math.MaxInt32
	for _, t := range tasks {
		ret = min(ret, t[0]+t[1])
	}
	return ret
}
