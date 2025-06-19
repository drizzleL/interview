package main

import (
	"math"
)

func assignEdgeWeights2(edges [][]int, queries [][]int) []int {
	n := len(edges) + 1
	ret := make([]int, len(queries))
	depth := make([]int, n)
	dict := make([][]int, n)
	for _, ed := range edges {
		dict[ed[0]-1] = append(dict[ed[0]-1], ed[1]-1)
		dict[ed[1]-1] = append(dict[ed[1]-1], ed[0]-1)
	}
	seen := make([]bool, n)
	var dfs func(i int, d int)
	var maxDepth int
	dirParent := make([]int, n)
	dirParent[0] = -1
	dfs = func(i int, d int) {
		maxDepth = max(maxDepth, d)
		seen[i] = true
		depth[i] = d
		for _, j := range dict[i] {
			if seen[j] {
				continue
			}
			dirParent[j] = i
			seen[j] = true
			dfs(j, d+1)
		}
	}
	dfs(0, 0)
	size := int(math.Ceil(math.Log2(float64(maxDepth))))
	parent := make([][]int, n)
	for i := range parent {
		parent[i] = make([]int, int(size)+1)
		for j := range parent[i] {
			parent[i][j] = -1
		}
	}
	for j := 0; j <= size; j++ {
		for i := 0; i < n; i++ {
			if j == 0 {
				parent[i][j] = dirParent[i]
				continue
			}
			if parent[i][j-1] == -1 {
				continue
			}
			parent[i][j] = parent[parent[i][j-1]][j-1]
		}
	}
	lca := func(a, b int) int {
		if depth[a] > depth[b] {
			a, b = b, a
		}
		d := depth[b] - depth[a]
		for d > 0 { //
			step := int(math.Log2(float64(d)))
			b = parent[b][step]
			d -= 1 << step
		}
		if a == b {
			return a
		}
		for d := size; d >= 0; d-- {
			if parent[a][d] != -1 && parent[a][d] != parent[b][d] {
				a, b = parent[a][d], parent[b][d]
			}
		}
		return parent[a][0]
	}
	for i, q := range queries {
		a, b := q[0]-1, q[1]-1
		if a == b {
			ret[i] = 0
			continue
		}
		r := lca(a, b)
		d := depth[a] + depth[b] - 2*depth[r]
		ret[i] = fastPow(2, d-1, 1)
	}
	return ret
}
