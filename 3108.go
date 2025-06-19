package main

func minimumCost3(n int, edges [][]int, query [][]int) []int {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] == x {
			return x
		}
		parent[x] = find(parent[x])
		return parent[x]
	}
	union := func(a, b int) {
		pa, pb := find(a), find(b)
		if pa > pb {
			pa, pb = pb, pa
		}
		parent[pb] = pa
	}
	for _, ed := range edges {
		union(ed[0], ed[1])
	}
	groupVal := make([]int, n)
	for i := range groupVal {
		groupVal[i] = -1
	}
	for _, ed := range edges {
		g := find(ed[0])
		if groupVal[g] == -1 {
			groupVal[g] = ed[2]
			continue
		}
		groupVal[g] &= ed[2]
	}
	ret := make([]int, len(query))
	for i, q := range query {
		g1, g2 := find(q[0]), find(q[1])
		if g1 != g2 {
			ret[i] = -1
			continue
		}
		ret[i] = groupVal[g1]
	}
	return ret
}
