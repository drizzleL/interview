package main

import "container/heap"

type MedianFinder struct {
	maxs IntHeap
	mins MaxIntHeap
}

/** initialize your data structure here. */
func MedianConstructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if this.mins.Len() != 0 && num < this.mins.IntHeap[0] {
		heap.Push(&this.mins, num)
	} else {
		heap.Push(&this.maxs, num)
	}
	if this.maxs.Len() > this.mins.Len()+1 {
		v := heap.Pop(&this.maxs)
		heap.Push(&this.mins, v)
	}
	if this.mins.Len() > this.maxs.Len() {
		v := heap.Pop(&this.mins)
		heap.Push(&this.maxs, v)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.maxs) == 0 {
		return 0
	}
	if this.maxs.Len() == this.mins.Len() {
		return (float64(this.maxs[0]) + float64(this.mins.IntHeap[0])) / 2
	}
	return float64(this.maxs[0])
}

type MaxIntHeap struct {
	IntHeap
}

func (h MaxIntHeap) Less(i, j int) bool { return h.IntHeap[i] > h.IntHeap[j] }
