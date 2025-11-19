package main

import (
	"container/heap"
)

func maximumProduct2(nums []int, k int) int {
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.(int) < b.(int)
		},
	}
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
	}
	for k != 0 {
		top := heap.Pop(h).(int)
		heap.Push(h, top+1)
		k -= 1
	}
	if h.Top().(int) == 0 {
		return 0
	}
	ret := 1
	for h.Len() != 0 {
		ret *= heap.Pop(h).(int)
	}
	return ret
}
