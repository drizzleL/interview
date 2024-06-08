package main

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] >= intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	var maxRight int
	var removed int
	for _, v := range intervals {
		if v[1] <= maxRight {
			removed += 1
		}
		maxRight = max(maxRight, v[1])
	}
	return len(intervals) - removed
}
