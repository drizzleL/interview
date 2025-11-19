package main

import (
	"container/heap"
	"sort"
)

func getSkyline(buildings [][]int) [][]int {
	var nodes [][2]int
	for _, bd := range buildings {
		l, r, h := bd[0], bd[1], bd[2]
		nodes = append(nodes, [2]int{l, -h}, [2]int{r, h})
	}
	sort.Slice(nodes, func(i, j int) bool {
		n1, n2 := nodes[i], nodes[j]
		if n1[0] == n2[0] {
			return n1[1] < n2[1]
		}
		return n1[0] < n2[0]
	})
	var ret [][]int
	var prevH int
	dict := map[int]int{}
	var h IntHeap
	getTop := func() int {
		for h.Len() > 0 {
			top := h[0]
			if dict[top] > 0 {
				dict[top] -= 1
				heap.Pop(&h)
				continue
			}
			return top
		}
		return 0
	}
	for _, node := range nodes {
		x, y := node[0], node[1]
		if y < 0 {
			heap.Push(&h, -y)
		} else {
			dict[y] += 1
		}
		currH := getTop()
		if prevH != currH {
			ret = append(ret, []int{x, currH})
			prevH = currH
		}
	}
	return ret
}
