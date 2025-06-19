package main

import "sort"

func countWays2(ranges [][]int) int {
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] == ranges[j][0] {
			return ranges[i][1] < ranges[j][1]
		}
		return ranges[i][0] < ranges[j][0]
	})
	pre := -1
	var groups int
	for i := 0; i < len(ranges); i++ {
		if ranges[i][0] > pre {
			groups += 1
		}
		pre = max(pre, ranges[i][1])
	}
	return fastPow(2, groups, 1)
}
