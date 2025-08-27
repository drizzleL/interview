package main

import "container/heap"

func lastStoneWeight(stones []int) int {
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.(int) > b.(int)
		},
	}
	for _, stone := range stones {
		heap.Push(h, stone)
	}
	for h.Len() > 1 {
		a, b := heap.Pop(h).(int), heap.Pop(h).(int)
		if a == b {
			continue
		}
		heap.Push(h, a-b)
	}
	if h.Len() == 0 {
		return 0
	}
	return h.Top().(int)
}
