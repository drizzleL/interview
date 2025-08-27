package main

import (
	"container/heap"
)

func eatenApples(apples []int, days []int) int {
	h := HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.([2]int)[1] < b.([2]int)[1]
		},
	}
	var ret int
	var i int
	for ; i < len(apples); i++ {
		if apples[i] != 0 {
			heap.Push(&h, [2]int{apples[i], days[i] + i})
		}
		for h.Len() != 0 {
			top := heap.Pop(&h).([2]int)
			if top[1] <= i { // remove expired
				continue
			}
			top[0] -= 1
			ret += 1
			if top[0] > 0 {
				heap.Push(&h, top)
			}
			break
		}
	}
	for h.Len() != 0 {
		top := heap.Pop(&h).([2]int)
		if top[1] <= i { // remove expired
			continue
		}
		tmp := min(top[0], top[1]-i)
		ret += tmp
		i += tmp
	}
	return ret
}

type HeapArr struct {
	Array
	LessHelper func(a, b interface{}) bool
}

type Array []interface{}

func (h Array) Top() interface{} { return h[0] }
func (h Array) Len() int         { return len(h) }
func (h HeapArr) Less(i, j int) bool {
	return h.LessHelper(h.Array[i], h.Array[j])
}
func (h Array) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *Array) Push(x any) {
	*h = append(*h, x)
}

func (h *Array) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
