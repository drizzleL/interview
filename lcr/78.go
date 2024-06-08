package main

import "container/heap"

func mergeKLists(lists []*ListNode) *ListNode {
	h := NodeHeap{}
	for _, l := range lists {
		if l == nil {
			continue
		}
		heap.Push(&h, l)
	}
	pre := &ListNode{}
	node := pre
	for h.Len() != 0 {
		n := heap.Pop(&h).(*ListNode)
		node.Next = n
		if n.Next != nil {
			heap.Push(&h, n.Next)
		}
		node = node.Next
	}
	return pre.Next
}

type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
