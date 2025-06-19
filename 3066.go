package main

import "container/heap"

func minOperations8(nums []int, k int) int {
	var h IntHeap
	var ret int
	for _, num := range nums {
		heap.Push(&h, num)
	}
	for h[0] < k {
		x := heap.Pop(&h).(int)
		y := heap.Pop(&h).(int)
		heap.Push(&h, min(x, y)*2+max(x, y))
		ret += 1
	}
	return ret
}
