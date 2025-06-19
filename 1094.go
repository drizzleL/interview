package main

import (
	"container/heap"
	"sort"
)

func carPooling(trips [][]int, capacity int) bool {
	sort.Slice(trips, func(i, j int) bool {
		return trips[i][1] < trips[j][1]
	})
	var cnt int
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.([]int)[2] < b.([]int)[2]
		},
	}
	for i := range trips {
		t := trips[i]
		for h.Len() != 0 {
			top := h.Array[0].([]int)
			if top[2] > t[1] {
				break
			}
			heap.Pop(h)
			cnt -= top[0]
		}
		cnt += t[0]
		if cnt > capacity {
			return false
		}
		heap.Push(h, t)
	}
	return true
}
