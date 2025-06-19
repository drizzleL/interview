package main

import (
	"math"
)

func minimumWeight(edges [][]int, queries [][]int) []int {
	n := len(edges) + 1
	dict := map[int][][2]int{}
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], [2]int{ed[1], ed[2]})
		dict[ed[1]] = append(dict[ed[1]], [2]int{ed[0], ed[2]})
	}
	seen := make([]bool, n)
	dis := make([]int, n)
	lvls := make([]int, n)
	directParent := make([]int, n)
	directParent[0] = -1
	var dfs func(x int, pre int, lvl int)
	var maxLvl int
	dfs = func(x int, pre int, lvl int) {
		maxLvl = max(maxLvl, lvl)
		dis[x] = pre
		seen[x] = true
		lvls[x] = lvl
		for _, next := range dict[x] {
			if seen[next[0]] {
				continue
			}
			directParent[next[0]] = x
			dfs(next[0], pre+next[1], lvl+1)
		}
	}
	dfs(0, 0, 0)
	size := int(math.Ceil(math.Log2(float64(maxLvl))))
	parents := make([][]int, n)
	for i := range parents {
		parents[i] = make([]int, size+1)
		for j := range parents[i] {
			parents[i][j] = -1
		}
	}
	for i := 0; i < n; i++ {
		parents[i][0] = directParent[i]
	}
	for j := 1; j <= size; j++ {
		for i := 0; i < n; i++ {
			if parents[i][j-1] == -1 {
				continue
			}
			parents[i][j] = parents[parents[i][j-1]][j-1]
		}
	}
	lca := func(a, b int) int {
		if lvls[a] > lvls[b] {
			a, b = b, a
		}
		d := lvls[b] - lvls[a]
		for d > 0 { //
			step := int(math.Log2(float64(d)))
			b = parents[b][step]
			d -= 1 << step
		}
		if a == b {
			return a
		}
		for d := size; d >= 0; d-- {
			if parents[a][d] != -1 && parents[a][d] != parents[b][d] {
				a, b = parents[a][d], parents[b][d]
			}
		}
		return parents[a][0]
	}
	getDis := func(a, b int) int {
		return dis[a] + dis[b] - 2*dis[lca(a, b)]
	}
	ret := make([]int, len(queries))
	for i, q := range queries {
		s1, s2, d := q[0], q[1], q[2]
		ret[i] = (getDis(s1, s2) + getDis(s1, d) + getDis(s2, d)) / 2
	}
	return ret
}
