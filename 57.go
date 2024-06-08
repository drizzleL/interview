package main

import (
	"sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	start := sort.Search(len(intervals), func(i int) bool {
		return intervals[i][1] >= newInterval[0]
	})
	end := sort.Search(len(intervals), func(i int) bool {
		return intervals[i][0] > newInterval[1]
	})
	if start != len(intervals) {
		newInterval[0] = min(newInterval[0], intervals[start][0])
	}
	if end != 0 {
		newInterval[1] = max(newInterval[1], intervals[end-1][1])
	}
	ret := make([][]int, start)
	copy(ret, intervals[:start])
	ret = append(ret, newInterval)
	ret = append(ret, intervals[end:]...)
	return ret
}
