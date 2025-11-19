package main

import (
	"container/heap"
	"math"
)

func minCost18(maxTime int, edges [][]int, passingFees []int) int {
	dict := make([][][2]int, len(passingFees))
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], [2]int{ed[1], ed[2]})
		dict[ed[1]] = append(dict[ed[1]], [2]int{ed[0], ed[2]})
	}
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			va, vb := a.([3]int), b.([3]int)
			return va[2] < vb[2]
		},
	}
	cache := make([]int, len(passingFees))
	for i := range cache {
		cache[i] = math.MaxInt32
	}
	cache[0] = 0
	// cache := make([][]int, len(passingFees))
	// for i := range cache {
	// 	cache[i] = make([]int, maxTime+1)
	// 	for j := range cache[i] {
	// 		cache[i][j] = math.MaxInt32
	// 	}
	// }
	// cache[0][0] = passingFees[0]
	heap.Push(h, [3]int{0, 0, passingFees[0]})
	for h.Len() > 0 {
		top := heap.Pop(h).([3]int)
		i, t, cost := top[0], top[1], top[2]
		if i == len(passingFees)-1 {
			return cost
		}
		for _, ed := range dict[i] {
			j, t2 := ed[0], ed[1]
			if t+t2 > maxTime {
				continue
			}
			newCost := cost + passingFees[j]
			if cache[j] <= t+t2 {
				continue
			}
			cache[j] = t + t2
			// cache[j][t+t2] = newCost
			heap.Push(h, [3]int{j, t + t2, newCost})
		}
	}
	return -1
}
