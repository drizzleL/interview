package main

import "container/heap"

type KthLargest struct {
	data IntHeap
	k    int
}

func ConstructorKth(k int, nums []int) KthLargest {
	ret := KthLargest{
		data: IntHeap{},
		k:    k,
	}
	for _, num := range nums {
		ret.Add(num)
	}
	return ret
}

func (this *KthLargest) Add(val int) int {
	if this.data.Len() == this.k && this.data[0] < val {
		heap.Pop(&this.data)
	}
	if this.data.Len() != this.k {
		heap.Push(&this.data, val)
	}
	return this.data[0]
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
