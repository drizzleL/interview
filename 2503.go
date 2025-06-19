package main

import (
	"sort"
)

func maxPoints2(grid [][]int, queries []int) []int {
	q2 := make([][2]int, 0, len(queries))
	for i, q := range queries {
		q2 = append(q2, [2]int{q, i})
	}
	sort.Slice(q2, func(i, j int) bool {
		return q2[i][0] < q2[j][0]
	})

	seen := make([][]bool, len(grid))
	for i := range seen {
		seen[i] = make([]bool, len(grid[0]))
	}
	roots := [][]int{{0, 0}}
	seen[0][0] = true
	var sum int
	var traverse func(nodes [][]int, limit int) [][]int
	traverse = func(nodes [][]int, limit int) [][]int {
		var ret [][]int
		for len(nodes) > 0 {
			var next [][]int
			for _, n := range nodes {
				x, y := n[0], n[1]
				if grid[x][y] >= limit { // still keeps
					ret = append(ret, n)
					continue
				}
				sum += 1
				for _, dir := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
					x2, y2 := x+dir[0], y+dir[1]
					if x2 < 0 || y2 < 0 || x2 >= len(grid) || y2 >= len(grid[0]) {
						continue
					}
					if seen[x2][y2] {
						continue
					}
					seen[x2][y2] = true
					next = append(next, []int{x2, y2})
				}
			}
			nodes = next
		}
		return ret
	}
	ret := make([]int, len(queries))
	for i := 0; i < len(q2); i++ {
		limit, idx := q2[i][0], q2[i][1]
		if i != 0 && q2[i][0] == q2[i-1][0] {
			ret[idx] = ret[q2[i-1][1]]
			continue
		}
		roots = traverse(roots, limit)
		ret[idx] = sum
	}
	return ret
}
