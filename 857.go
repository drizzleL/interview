package main

import (
	"container/heap"
	"math"
	"sort"
)

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	var eles []*Ele
	for i := range quality {
		eles = append(eles, &Ele{
			Q: float64(quality[i]),
			R: float64(wage[i]) / float64(quality[i]),
		})
	}
	sort.Slice(eles, func(i, j int) bool {
		return eles[i].R < eles[j].R
	})
	var h EleHeap
	var sum float64
	topR := eles[k-1].R
	for i := 0; i < k; i++ {
		sum += float64(eles[i].Q)
		heap.Push(&h, eles[i])
	}
	ret := sum * topR
	for i := k; i < len(eles); i++ {
		top := heap.Pop(&h).(*Ele)
		sum -= top.Q
		sum += eles[i].Q
		heap.Push(&h, eles[i])
		topR = eles[i].R
		ret = math.Min(ret, sum*topR)
	}
	return ret
}

type Ele struct {
	Q float64
	R float64
}

type EleHeap []*Ele

func (h EleHeap) Len() int           { return len(h) }
func (h EleHeap) Less(i, j int) bool { return h[i].Q > h[j].Q }
func (h EleHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EleHeap) Push(x any) {
	*h = append(*h, x.(*Ele))
}

func (h *EleHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
