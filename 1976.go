package main

import (
	"container/heap"
	"math"
)

func countPaths2(n int, roads [][]int) int {
	dict := map[int][][]int{}
	for _, r := range roads {
		dict[r[0]] = append(dict[r[0]], []int{r[1], r[2]})
		dict[r[1]] = append(dict[r[1]], []int{r[0], r[2]})
	}
	h := HeapArr{
		LessHelper: func(a, b interface{}) bool {
			v1, v2 := a.([2]int), b.([2]int)
			return v1[1] < v2[1]
		},
	}
	time := make([]int, n)
	for i := range time {
		time[i] = math.MaxInt64
	}
	paths := make([]int, n)
	paths[0] = 1
	time[0] = 0
	heap.Push(&h, [2]int{0, 0})
	for h.Len() != 0 {
		top := heap.Pop(&h).([2]int)
		i, t := top[0], top[1]
		if i == n-1 {
			return paths[i]
		}
		if time[i] != t { // not match, skip
			continue
		}
		for _, way := range dict[i] {
			child, t2 := way[0], way[1]
			if t+t2 > time[child] {
				continue
			}
			if t+t2 < time[child] {
				time[child] = t + t2
				paths[child] = 0
				heap.Push(&h, [2]int{child, t + t2})
			}
			paths[child] += paths[i]
			paths[child] %= 1e9 + 7
		}
	}
	return 0
}
