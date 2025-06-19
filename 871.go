package main

import (
	"container/heap"
	"sort"
)

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	var ret int
	sort.Slice(stations, func(i, j int) bool {
		return stations[i][0] < stations[j][0]
	})
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.(int) > b.(int)
		},
	}
	for pos, j := startFuel, 0; pos < target; {
		for ; j < len(stations) && stations[j][0] <= pos; j++ {
			heap.Push(h, stations[j][1])
		}
		if h.Len() == 0 {
			return -1
		}
		pos += heap.Pop(h).(int)
		ret += 1
	}
	return ret
}
