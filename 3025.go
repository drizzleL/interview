package main

import (
	"math"
	"sort"
)

func numberOfPairs3(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	var ret int
	for i := 0; i < len(points); i++ {
		y := math.MinInt32
		for j := i + 1; j < len(points); j++ {
			if points[j][1] > points[i][1] {
				continue
			}
			if points[j][1] >= y {
				y = points[j][1]
				ret += 1
			}
		}
	}
	return ret
}
