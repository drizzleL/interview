package main

import (
	"container/heap"
)

type SafeEle struct {
	I, J int
	Safe int
}

type SafeEles []*SafeEle

func (se SafeEles) Len() int {
	return len(se)
}
func (se SafeEles) Swap(i, j int) {
	se[i], se[j] = se[j], se[i]
}

func (se SafeEles) Less(i, j int) bool {
	return se[i].Safe > se[j].Safe
}

func (se *SafeEles) Push(v interface{}) {
	*se = append(*se, v.(*SafeEle))
}
func (se *SafeEles) Pop() interface{} {
	old := *se
	val := old[len(old)-1]
	*se = old[0 : len(old)-1]
	return val
}

func maximumSafenessFactor(grid [][]int) int {
	size := len(grid)
	if size == 0 {
		return 0
	}
	if grid[0][0] == 1 || grid[size-1][size-1] == 1 {
		return 0
	}
	var nodes [][2]int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[i][j] == 1 {
				nodes = append(nodes, [2]int{i, j})
			}
		}
	}
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for step := 2; len(nodes) != 0; step++ {
		var next [][2]int
		for _, node := range nodes {
			for _, dir := range dirs {
				i2, j2 := node[0]+dir[0], node[1]+dir[1]
				if i2 < 0 || j2 < 0 || i2 >= size || j2 >= size {
					continue
				}
				if grid[i2][j2] != 0 {
					continue
				}
				grid[i2][j2] = step
				next = append(next, [2]int{i2, j2})
			}
		}
		nodes = next
	}
	var q SafeEles
	seen := make([][]bool, size)
	for i := range seen {
		seen[i] = make([]bool, size)
	}
	seen[0][0] = true
	heap.Push(&q, &SafeEle{
		I:    0,
		J:    0,
		Safe: grid[0][0],
	})
	for q.Len() != 0 {
		ele := heap.Pop(&q).(*SafeEle)
		for _, dir := range dirs {
			i2, j2 := ele.I+dir[0], ele.J+dir[1]
			if i2 < 0 || j2 < 0 || i2 >= size || j2 >= size {
				continue
			}
			if grid[i2][j2] == 1 {
				continue
			}
			if seen[i2][j2] {
				continue
			}
			seen[i2][j2] = true
			if i2 == size-1 && j2 == size-1 {
				return min(grid[i2][j2], ele.Safe) - 1
			}
			heap.Push(&q, &SafeEle{
				I:    i2,
				J:    j2,
				Safe: min(grid[i2][j2], ele.Safe),
			})
		}
	}
	return 0
}
