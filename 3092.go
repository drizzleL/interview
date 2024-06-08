package main

import "container/heap"

func mostFrequentIDs(nums []int, freq []int) []int64 {
	dict := map[int]*FreqItem{}
	var h FreqItems
	ret := make([]int64, len(nums))
	for i := range nums {
		n, f := nums[i], freq[i]
		if old, ok := dict[n]; ok {
			f += old.Freq
			old.Deprecated = true
		}
		item := &FreqItem{
			Num:  n,
			Freq: f,
		}
		dict[n] = item
		heap.Push(&h, item)
		for len(h) != 0 && h[0].Deprecated {
			heap.Pop(&h)
		}
		ret[i] = int64(h[0].Freq)
	}
	return ret
}

type FreqItem struct {
	Num        int
	Freq       int
	Deprecated bool
}

type FreqItems []*FreqItem

func (h FreqItems) Len() int           { return len(h) }
func (h FreqItems) Less(i, j int) bool { return h[i].Freq > h[j].Freq }
func (h FreqItems) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *FreqItems) Push(x any) {
	*h = append(*h, x.(*FreqItem))
}

func (h *FreqItems) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
