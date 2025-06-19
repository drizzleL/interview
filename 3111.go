package main

import "sort"

func minRectanglesToCoverPoints(points [][]int, w int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	var ret int
	for i := 0; i < len(points); {
		ret += 1
		start := points[i][0]
		end := start + w
		for ; i < len(points) && points[i][0] <= end; i++ {
		}
	}
	return ret
}
