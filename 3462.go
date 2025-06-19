package main

import (
	"container/heap"
	"sort"
)

func maxSum2(grid [][]int, limits []int, k int) int64 {
	m, n := len(grid), len(grid[0])
	h := HeapArr{
		LessHelper: func(a, b interface{}) bool {
			ele1, ele2 := a.([2]int), b.([2]int)
			return grid[ele1[0]][ele1[1]] > grid[ele2[0]][ele2[1]]
		},
	}
	for i := 0; i < m; i++ {
		sort.Ints(grid[i])
		if limits[i] != 0 {
			heap.Push(&h, [2]int{i, n - 1})
		}
	}
	var ret int
	for k > 0 {
		ele := heap.Pop(&h).([2]int)
		i, j := ele[0], ele[1]
		ret += grid[i][j]
		limits[i] -= 1
		if limits[i] != 0 {
			heap.Push(&h, [2]int{i, j - 1})
		}
		k -= 1
	}
	return int64(ret)
}
