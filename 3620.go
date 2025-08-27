package main

import (
	"container/heap"
	"math"
)

func findMaxPathScore(edges [][]int, online []bool, k int64) int {
	dict := make([][]int, len(online))
	var l, r int
	for i, ed := range edges {
		if !online[ed[0]] || !online[ed[1]] {
			continue
		}
		r = max(r, ed[2])
		dict[ed[0]] = append(dict[ed[0]], i)
	}
	check := func(x int) bool {
		h := &HeapArr{
			LessHelper: func(a, b interface{}) bool {
				return a.([3]int)[0] < b.([3]int)[0]
			},
		}
		seen := make([]int, len(online))
		heap.Push(h, [3]int{0, 0, math.MaxInt32})
		for h.Len() > 0 {
			top := heap.Pop(h).([3]int)
			if top[2] <= seen[top[1]] {
				continue
			}
			if top[1] == len(online)-1 {
				return true
			}
			seen[top[1]] = top[2]
			for _, next := range dict[top[1]] {
				ed := edges[next]
				if ed[2] < x {
					continue
				}
				total := top[0] + ed[2]
				if total > int(k) {
					continue
				}
				heap.Push(h, [3]int{total, ed[1], min(top[2], ed[2])})
			}
		}
		return false
	}
	if !check(0) {
		return -1
	}
	for l < r {
		mid := l + (r-l+1)/2
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}
