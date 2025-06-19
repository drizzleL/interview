package main

import "container/heap"

func minStoneSum(piles []int, k int) int {
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.(int) > b.(int)
		},
	}
	var ret int
	for _, p := range piles {
		heap.Push(h, p)
		ret += p
	}
	for i := 0; i < k; i++ {
		top := heap.Pop(h).(int)
		if top == 1 { // no meaning
			break
		}
		tmp := top / 2
		ret -= tmp
		heap.Push(h, top-tmp)
	}
	return ret
}
