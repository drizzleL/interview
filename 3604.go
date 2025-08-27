package main

import "container/heap"

func minTime5(n int, edges [][]int) int {
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.([2]int)[0] < b.([2]int)[0]
		},
	}
	dict := make([][]int, n)
	for i, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], i)
	}
	seen := make([]bool, n)
	seen[0] = true
	heap.Push(h, [2]int{0, 0})
	for h.Len() > 0 {
		top := heap.Pop(h).([2]int)
		if seen[top[1]] {
			continue
		}
		if top[1] == n-1 {
			return top[0]
		}
		seen[top[1]] = true
		for _, next := range dict[top[1]] {
			ed := edges[next]
			if seen[ed[1]] {
				continue
			}
			if top[0]+1 > ed[3] { // exceed time
				continue
			}
			heap.Push(h, [2]int{max(top[0]+1, ed[2]), ed[1]})
		}
	}
	return -1
}
