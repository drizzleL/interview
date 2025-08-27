package main

import (
	"container/heap"
	"math"
)

func minCost17(n int, edges [][]int) int {
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.([2]int)[0] < b.([2]int)[0]
		},
	}
	dict := make([][][2]int, n)
	revDict := make([][][2]int, n)
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], [2]int{ed[1], ed[2]})
		revDict[ed[1]] = append(revDict[ed[1]], [2]int{ed[0], ed[2] * 2})
	}
	heap.Push(h, [2]int{0, 0})
	seen := make([]int, n)
	for i := range seen {
		seen[i] = math.MaxInt32
	}
	for h.Len() > 0 {
		top := heap.Pop(h).([3]int)
		cost, i := top[0], top[1]
		if seen[i] <= cost {
			continue
		}
		if i == n-1 {
			return cost
		}
		for _, next := range dict[i] {
			nextCost := cost + next[1]
			if seen[next[0]] <= nextCost {
				continue
			}
			seen[next[0]] = nextCost
			heap.Push(h, [2]int{nextCost, next[0]})
		}
		for _, next := range revDict[i] {
			nextCost := cost + next[1]
			if seen[next[0]] <= nextCost {
				continue
			}
			seen[next[0]] = nextCost
			heap.Push(h, [2]int{nextCost, next[0]})
		}
	}
	return -1
}
