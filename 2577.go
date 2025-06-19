package main

import (
	"container/heap"
)

func minimumTime2(grid [][]int) int {
	if grid[0][0] != 0 {
		return -1
	}
	m, n := len(grid), len(grid[0])
	if grid[0][1] > 1 && grid[1][0] > 1 {
		return -1
	}
	var h GridNodes
	heap.Push(&h, &GridNode{
		Step:  0,
		Coord: [2]int{0, 0},
	})
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	seen[0][0] = true
	for h.Len() != 0 {
		top := heap.Pop(&h).(*GridNode)
		x, y := top.Coord[0], top.Coord[1]
		if x == m-1 && y == n-1 {
			return top.Step
		}
		for _, dir := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			x2, y2 := x+dir[0], y+dir[1]
			if x2 < 0 || y2 < 0 || x2 >= m || y2 >= n { // out of bounds
				continue
			}
			if seen[x2][y2] { // no back
				continue
			}
			seen[x2][y2] = true
			step := top.Step + 1
			if grid[x2][y2] > step {
				step += (grid[x2][y2] - step + 1) / 2 * 2
			}
			heap.Push(&h, &GridNode{
				Coord: [2]int{x2, y2},
				Step:  step,
			})
		}
	}
	return 0
}

type GridNode struct {
	Coord [2]int
	Step  int
}

type GridNodes []*GridNode

func (h GridNodes) Len() int           { return len(h) }
func (h GridNodes) Less(i, j int) bool { return h[i].Step < h[j].Step }
func (h GridNodes) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *GridNodes) Push(x any) {
	*h = append(*h, x.(*GridNode))
}

func (h *GridNodes) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
