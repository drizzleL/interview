package main

import (
	"sort"
)

func numberOfGoodPaths(vals []int, edges [][]int) int {
	parent := make([]int, len(vals))
	for i := range parent {
		parent[i] = i
	}
	var find func(i int) int
	find = func(i int) int {
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return parent[i]
	}
	union := func(a, b int) {
		pa, pb := find(a), find(b)
		if pa > pb {
			pa, pb = pb, pa
		}
		parent[pb] = pa
	}
	dict := make([][]int, len(vals))
	for _, ed := range edges {
		sort.Slice(ed, func(i, j int) bool {
			return vals[ed[i]] < vals[ed[j]]
		})
		dict[ed[1]] = append(dict[ed[1]], ed[0])
	}
	var pairs [][2]int
	for i, v := range vals {
		pairs = append(pairs, [2]int{i, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	var ret int
	for i := 0; i < len(pairs); {
		pivot := pairs[i][1]
		var cand []int
		for i < len(pairs) && pairs[i][1] == pivot {
			p := pairs[i]
			cand = append(cand, p[0])
			for _, child := range dict[p[0]] {
				union(p[0], child)
			}
			i++
		}
		dict := map[int]int{}
		for _, c := range cand {
			dict[find(c)] += 1
		}
		for _, v := range dict {
			ret += v*(v-1)/2 + v
		}
	}
	return ret
}
