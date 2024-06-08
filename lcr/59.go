package main

import "container/heap"

type KthLargest struct {
	h    *IntHeap
	size int
}

func KthConstructor(k int, nums []int) KthLargest {
	h := IntHeap(nil)
	l := KthLargest{
		h:    &h,
		size: k,
	}
	for _, num := range nums {
		l.Add(num)
	}
	return l
}

func (this *KthLargest) Add(val int) int {
	if this.h.Len() != this.size {
		heap.Push(this.h, val)
		return (*this.h)[0]
	}
	if val > (*this.h)[0] {
		heap.Pop(this.h)
		heap.Push(this.h, val)
	}
	return (*this.h)[0]
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
