package main

import "container/heap"

func maxKelements(nums []int, k int) int64 {
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.(int) > b.(int)
		},
	}
	for _, num := range nums {
		heap.Push(h, num)
	}
	var ret int
	for i := 0; i < k; i++ {
		top := heap.Pop(h).(int)
		ret += top
		heap.Push(h, (top+2)/3)
	}
	return int64(ret)
}
