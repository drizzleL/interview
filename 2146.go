package main

import "container/heap"

func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int {
	if grid[start[0]][start[1]] == 0 {
		return nil
	}
	h := &HeapArr{
		LessHelper: func(a, b interface{}) bool {
			av, bv := a.([4]int), b.([4]int)
			if av[3] != bv[3] {
				return av[3] < bv[3]
			}
			if av[2] != bv[2] {
				return av[2] < bv[2]
			}
			if av[0] != bv[0] {
				return av[0] < bv[0]
			}
			return av[1] < bv[1]
		},
	}
	seen := make([][]bool, len(grid))
	for i := range seen {
		seen[i] = make([]bool, len(grid[0]))
	}
	var ret [][]int
	nodes := [][]int{start}
	for step := 0; len(nodes) > 0 && len(ret) < k; step++ {
		var next [][]int
		for _, node := range nodes {
			seen[node[0]][node[1]] = true
			p := grid[node[0]][node[1]]
			if p >= pricing[0] && p <= pricing[1] {
				heap.Push(h, [4]int{node[0], node[1], p, step})
			}
			for _, dir := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				x2, y2 := node[0]+dir[0], node[1]+dir[1]
				if x2 < 0 || y2 < 0 || x2 >= len(grid) || y2 >= len(grid[0]) {
					continue
				}
				if seen[x2][y2] {
					continue
				}
				if grid[x2][y2] == 0 {
					continue
				}
				seen[x2][y2] = true
				next = append(next, []int{x2, y2})
			}
		}
		for len(ret) < k && h.Len() != 0 {
			top := heap.Pop(h).([4]int)
			ret = append(ret, []int{top[0], top[1]})
		}
		nodes = next
	}
	return ret
}
