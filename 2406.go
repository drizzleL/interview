package main

import (
	"container/heap"
	"sort"
)

func minGroups(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	var h IntHeap
	for _, intv := range intervals {
		if h.Len() != 0 && h[0] >= intv[0] {
			heap.Pop(&h)
		}
		heap.Push(&h, intv[1])
	}
	return h.Len()
}
