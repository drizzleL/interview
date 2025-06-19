package main

import (
	"container/heap"
)

func minimumTime3(n int, edges [][]int, disappear []int) []int {
	h := HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.([2]int)[1] < b.([2]int)[1]
		},
	}
	dict := map[int][][2]int{}
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], [2]int{ed[1], ed[2]})
		dict[ed[1]] = append(dict[ed[1]], [2]int{ed[0], ed[2]})
	}
	ret := make([]int, n)
	for i := range ret {
		ret[i] = -1
	}
	heap.Push(&h, [2]int{0, 0})
	for h.Len() != 0 {
		top := heap.Pop(&h).([2]int)
		idx, takes := top[0], top[1]
		if ret[idx] != -1 {
			continue
		}
		ret[idx] = takes
		for _, child := range dict[idx] {
			idx2, takes2 := child[0], child[1]
			if takes+takes2 >= disappear[idx2] {
				continue
			}
			heap.Push(&h, [2]int{idx2, takes + takes2})
		}
	}
	return ret
}
