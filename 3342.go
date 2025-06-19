package main

import "container/heap"

func minTimeToReach2(moveTime [][]int) int {
	if len(moveTime) == 0 {
		return 0
	}
	m, n := len(moveTime), len(moveTime[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var h GridHeap
	heap.Push(&h, &Grid{
		Step:  0,
		X:     0,
		Y:     0,
		Extra: 1,
	})
	visited[0][0] = true
	for h.Len() != 0 {
		g := heap.Pop(&h).(*Grid)
		if g.X == m-1 && g.Y == n-1 {
			return g.Step
		}
		for _, dir := range [][4]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
			x, y := g.X+dir[0], g.Y+dir[1]
			if x < 0 || y < 0 || x >= m || y >= n {
				continue
			}
			if visited[x][y] {
				continue
			}
			visited[x][y] = true
			heap.Push(&h, &Grid{
				Step:  max(moveTime[x][y], g.Step) + g.Extra,
				X:     x,
				Y:     y,
				Extra: 3 - g.Extra,
			})
		}
	}
	return 0
}
