package main

func maxNumEdgesToRemove(n int, edges [][]int) int {
	parents := make([]int, n+1)
	for i := range parents {
		parents[i] = i
	}
	var find func(i int) int
	find = func(i int) int {
		if parents[i] != i {
			parents[i] = find(parents[i])
		}
		return parents[i]
	}
	union := func(a, b int) bool {
		pa, pb := find(a), find(b)
		if pa == pb {
			return false
		}
		if pa > pb {
			pa, pb = pb, pa
		}
		parents[pb] = pa
		return true
	}
	var ret, ed1, ed2 int
	for _, ed := range edges {
		if ed[0] != 3 {
			continue
		}
		if union(ed[1], ed[2]) {
			ed1 += 1
			ed2 += 1
		} else {
			ret += 1
		}
	}
	cp := make([]int, n+1)
	copy(cp, parents)
	for _, ed := range edges {
		if ed[0] != 1 {
			continue
		}
		if union(ed[1], ed[2]) {
			ed2 += 1
		} else {
			ret += 1
		}
	}
	parents = cp
	for _, ed := range edges {
		if ed[0] != 2 {
			continue
		}
		if union(ed[1], ed[2]) {
			ed1 += 1
		} else {
			ret += 1
		}
	}
	if ed1 == ed2 && ed1 == n-1 {
		return ret
	}
	return -1
}
