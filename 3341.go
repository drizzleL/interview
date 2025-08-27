package main

import (
	"container/heap"
	"math"
)

func minTimeToReach(moveTime [][]int) int {
	m, n := len(moveTime), len(moveTime[0])
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.([3]int)[0] < b.([3]int)[0]
		},
	}
	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
		for j := range visited[i] {
			visited[i][j] = math.MaxInt32
		}
	}
	visited[0][0] = 0
	heap.Push(h, [3]int{0, 0, 0})
	for h.Len() != 0 {
		top := heap.Pop(h).([3]int)
		if top[1] == m-1 && top[2] == n-1 {
			return top[0]
		}
		for _, dir := range [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
			x2, y2 := top[1]+dir[0], top[2]+dir[1]
			if x2 < 0 || y2 < 0 || x2 >= m || y2 >= n {
				continue
			}
			if visited[x2][y2] <= top[0]+1 {
				continue
			}
			cost := max(moveTime[x2][y2]+1, top[0]+1)
			visited[x2][y2] = cost
			heap.Push(h, [3]int{cost, x2, y2})
		}
	}
	return 0
}
