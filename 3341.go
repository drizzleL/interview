package main

import "container/heap"

func minTimeToReach(moveTime [][]int) int {
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
		Step: max(moveTime[0][1], 0),
		X:    0,
		Y:    1,
	})
	heap.Push(&h, &Grid{
		Step: max(moveTime[1][0], 0),
		X:    1,
		Y:    0,
	})
	visited[0][1] = true
	visited[1][0] = true
	for h.Len() != 0 {
		g := heap.Pop(&h).(*Grid)
		if g.X == m-1 && g.Y == n-1 {
			return g.Step + 1
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
				Step: max(moveTime[x][y], g.Step+1),
				X:    x,
				Y:    y,
			})
		}
	}
	return 0
}

type GridHeap []*Grid

type Grid struct {
	Step  int
	X     int
	Y     int
	Extra int
}

func (h GridHeap) Len() int           { return len(h) }
func (h GridHeap) Less(i, j int) bool { return h[i].Step < h[j].Step }
func (h GridHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *GridHeap) Push(x any) {
	*h = append(*h, x.(*Grid))
}

func (h *GridHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
