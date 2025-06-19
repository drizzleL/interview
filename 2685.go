package main

func countCompleteComponents(n int, edges [][]int) int {
	parent := make([]int, n)
	out := make([]int, n)
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
	for _, ed := range edges {
		union(ed[0], ed[1])
		out[ed[0]] += 1
		out[ed[1]] += 1
	}
	group := make([][]int, n)
	for i := range parent {
		g := find(i)
		group[g] = append(group[g], i)
	}
	var ret int
	for _, v := range group {
		if len(v) == 0 {
			continue
		}
		var flag bool
		for _, child := range v {
			if out[child] != len(v)-1 {
				flag = true
				break
			}
		}
		if !flag {
			ret += 1
		}
	}
	return ret
}
