package main

import (
	"container/heap"
	"math"
	"sort"
)

func maxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] == events[j][0] {
			return events[i][1] < events[j][1]
		}
		return events[i][0] < events[j][0]
	})
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.([]int)[1] < b.([]int)[1]
		},
	}
	var ret int
	last := math.MinInt32
	for i := 0; i < len(events) || h.Len() != 0; {
		for i < len(events) && events[i][0] <= last {
			heap.Push(h, events[i])
			i++
		}
		if h.Len() == 0 {
			last = events[i][0] + 1
			ret += 1
			i++
			continue
		}
		top := heap.Pop(h).([]int)
		if last > top[1] {
			continue
		}
		ret += 1
		last += 1
	}
	return ret
}
