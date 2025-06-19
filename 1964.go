package main

import "sort"

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	var q []int
	ret := make([]int, len(obstacles))
	for i, h := range obstacles {
		idx := sort.SearchInts(q, h+1)
		if idx == len(q) {
			q = append(q, h)
			ret[i] = len(q)
			continue
		}
		q[idx] = h
		ret[i] = idx + 1
	}
	return ret
}
