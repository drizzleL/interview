package main

import (
	"math"
	"sort"
)

func maxStability(n int, edges [][]int, k int) int {
	sort.Slice(edges, func(i, j int) bool {
		if edges[i][3] == edges[j][3] {
			return edges[i][2] > edges[j][2]
		}
		return edges[i][3] > edges[j][3]
	})
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(a, b int) {
		pa, pb := find(a), find(b)
		if pa > pb {
			pa, pb = pb, pa
		}
		parent[pb] = pa
	}
	ret := math.MaxInt32
	var s []int
	redundant := func(i, j int) bool {
		return find(i) == find(j)
	}
	for _, ed := range edges {
		if ed[3] == 1 { // must, just add
			ret = min(ret, ed[2])
			if redundant(ed[0], ed[1]) {
				return -1
			}
			union(ed[0], ed[1])
			continue
		}
		if redundant(ed[0], ed[1]) { // continue
			continue
		}
		s = append(s, ed[2])
		union(ed[0], ed[1])
	}
	for i := range parent {
		if find(i) != 0 {
			return -1
		}
	}
	sort.Ints(s)
	for i := 0; i < len(s); i++ {
		if i < k {
			ret = min(ret, s[i]*2)
		} else {
			ret = min(ret, s[i])
		}
	}
	return ret
}
