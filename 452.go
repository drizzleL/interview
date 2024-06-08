package main

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		p1, p2 := points[i], points[j]
		if p1[0] == p2[0] {
			return p1[1] < p2[1]
		}
		return p1[0] < p2[0]
	})
	var ret int
	lastInterval := []int{points[0][0] - 1, points[0][0] - 1}
	for _, p := range points {
		if p[0] > lastInterval[1] {
			lastInterval = p
			ret += 1
			continue
		}
		lastInterval[0] = max(lastInterval[0], p[0])
		lastInterval[1] = min(lastInterval[1], p[1])
	}
	return ret
}
