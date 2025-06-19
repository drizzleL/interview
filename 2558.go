package main

import (
	"container/heap"
	"math"
)

func pickGifts(gifts []int, k int) int64 {
	q := HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.(int) > b.(int)
		},
	}
	var sum int
	for _, g := range gifts {
		sum += g
		heap.Push(&q, g)
	}
	for i := 0; i < k; i++ {
		top := heap.Pop(&q).(int)
		left := int(math.Sqrt(float64(top)))
		take := top - left
		if take == 0 {
			break
		}
		sum -= take
		heap.Push(&q, left)
	}
	return int64(sum)
}
