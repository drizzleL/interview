package main

import (
	"container/heap"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	ret := [][]int{}
	h := IntsHeap{}
	for i := 0; i < len(nums1); i++ {
		heap.Push(&h, &Item{
			Idx0: i,
			Idx1: 0,
			Val:  nums1[i] + nums2[0],
		})
	}
	for len(ret) != k && h.Len() != 0 {
		top := heap.Pop(&h)
		item := top.(*Item)
		ret = append(ret, []int{nums1[item.Idx0], nums2[item.Idx1]})
		if item.Idx1 != len(nums2)-1 {
			item.Idx1++
			item.Val = nums1[item.Idx0] + nums2[item.Idx1]
			heap.Push(&h, item)
		}
	}
	return ret
}

type Item struct {
	Idx0, Idx1 int
	Val        int
}
type IntsHeap []*Item

func (h IntsHeap) Len() int           { return len(h) }
func (h IntsHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h IntsHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntsHeap) Push(x any) {
	*h = append(*h, x.(*Item))
}

func (h *IntsHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
