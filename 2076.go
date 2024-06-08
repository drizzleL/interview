package main

func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	parents := make([]int, n)
	for i := range parents {
		parents[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parents[x] == x {
			return x
		}
		parents[x] = find(parents[x])
		return parents[x]
	}
	union := func(a, b int) {
		pa, pb := find(a), find(b)
		if pa == pb {
			return
		}
		if pa > pb {
			pa, pb = pb, pa
		}
		parents[pb] = pa
	}
	ret := make([]bool, len(requests))
	for i, req := range requests {
		a, b := req[0], req[1]
		pa, pb := find(a), find(b)
		if find(a) == find(b) { // already connected
			ret[i] = true
			continue
		}
		var flag bool
		for _, restri := range restrictions {
			p, q := find(restri[0]), find(restri[1])
			if (p == pa && q == pb) || (p == pb && q == pa) {
				flag = true
				break
			}
		}
		if !flag {
			union(a, b)
			ret[i] = true
		}
	}
	return ret
}
