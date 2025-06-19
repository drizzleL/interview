package main

import "container/heap"

func furthestBuilding(heights []int, bricks int, ladders int) int {
	var h IntHeap
	for i := 1; i < len(heights); i++ {
		if heights[i] <= heights[i-1] {
			continue
		}
		gap := heights[i] - heights[i-1]
		if h.Len() < ladders {
			heap.Push(&h, gap)
			continue
		}
		if h.Len() != 0 && gap > h[0] {
			heap.Push(&h, gap)
			gap = heap.Pop(&h).(int)
		}
		if bricks < gap {
			return i - 1
		}
		bricks -= gap
	}
	return len(heights) - 1
}
