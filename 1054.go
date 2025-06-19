package main

import "container/heap"

func rearrangeBarcodes(barcodes []int) []int {
	dict := map[int]int{}
	for _, bar := range barcodes {
		dict[bar] += 1
	}
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.([2]int)[1] > b.([2]int)[1]
		},
	}
	for k, v := range dict {
		heap.Push(h, [2]int{k, v})
	}
	var ret []int
	for h.Len() != 0 {
		top := heap.Pop(h).([2]int)
		ret = append(ret, top[0])
		if h.Len() == 0 {
			break
		}
		sec := heap.Pop(h).([2]int)
		ret = append(ret, sec[0])
		top[1] -= 1
		if top[1] != 0 {
			heap.Push(h, top)
		}
		sec[1] -= 1
		if sec[1] != 0 {
			heap.Push(h, sec)
		}
	}
	return ret
}
