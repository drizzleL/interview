package main

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	if len(intervals) == 0 || newInterval[0] > intervals[len(intervals)-1][1] {
		return append(intervals, newInterval)
	}
	if newInterval[1] < intervals[0][0] {
		return append([][]int{newInterval}, intervals...)
	}
	var ret [][]int
	for len(intervals) != 0 {
		intv := intervals[0]
		intervals = intervals[1:]
		if intv[1] < newInterval[0] {
			ret = append(ret, intv)
			continue
		}
		if newInterval[1] < intv[0] {
			ret = append(ret, newInterval, intv)
			ret = append(ret, intervals...)
			return ret
		}
		ret = append(ret, []int{min(intv[0], newInterval[0]), max(intv[1], newInterval[1])})
		break
	}
	for len(intervals) != 0 {
		intv := intervals[0]
		if intv[0] > ret[len(ret)-1][1] {
			return append(ret, intervals...)
		}
		ret[len(ret)-1][1] = max(ret[len(ret)-1][1], intv[1])
		intervals = intervals[1:]
	}
	return ret
}
