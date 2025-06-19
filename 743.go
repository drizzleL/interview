package main

import "container/heap"

func networkDelayTime(times [][]int, n int, k int) int {
	dict := make([][][2]int, n)
	for _, t := range times {
		dict[t[0]-1] = append(dict[t[0]-1], [2]int{t[1] - 1, t[2]})
	}
	seen := make([]int, n)
	for i := range seen {
		seen[i] = -1
	}
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.([2]int)[1] < b.([2]int)[1]
		},
	}
	heap.Push(h, [2]int{k - 1, 0})
	for h.Len() > 0 {
		top := heap.Pop(h).([2]int)
		if seen[top[0]] != -1 {
			continue
		}
		seen[top[0]] = top[1]
		for _, next := range dict[top[0]] {
			if seen[next[0]] != -1 && seen[next[0]] < top[1]+next[1] {
				continue
			}
			heap.Push(h, [2]int{next[0], top[1] + next[1]})
		}
	}
	var ret int
	for _, v := range seen {
		if v == -1 {
			return -1
		}
		ret = max(ret, v)
	}
	return ret
}
