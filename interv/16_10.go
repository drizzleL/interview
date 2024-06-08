package main

import (
	"container/heap"
	"sort"
)

func maxAliveYear(birth []int, death []int) int {
	type dataTuple struct {
		b int
		d int
	}
	data := []dataTuple{}
	for i := 0; i < len(birth); i++ {
		data = append(data, dataTuple{birth[i], death[i]})
	}
	sort.Slice(data, func(i, j int) bool {
		if data[i].b == data[j].b {
			return data[i].d < data[j].d
		}
		return data[i].b < data[j].b
	})

	h := IntHeap{}
	var retCnt, year int
	for i := 0; i < len(birth); i++ {
		b := data[i].b
		for len(h) != 0 && h[0] < b {
			heap.Pop(&h)
		}
		heap.Push(&h, data[i].d)
		if len(h) > retCnt {
			retCnt = len(h)
			year = b
		}
	}
	return year
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
