package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	tmp := intervals[0]
	var ret [][]int
	for i := 1; i < len(intervals); i++ {
		intv := intervals[i]
		if intv[0] > tmp[1] {
			ret = append(ret, tmp)
			tmp = intv
			continue
		}
		if intv[1] >= tmp[1] {
			tmp[1] = intv[1]
		}
	}
	ret = append(ret, tmp)
	return ret
}
