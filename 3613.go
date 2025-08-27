package main

func minCost99(n int, edges [][]int, k int) int {
	var l, r int
	for _, ed := range edges {
		r = max(r, ed[2])
	}
	type uf struct {
		parent []int
		conn   int
	}
	newUf := func(n int) *uf {
		parent := make([]int, n)
		for i := range parent {
			parent[i] = i
		}
		u := &uf{
			conn:   n,
			parent: parent,
		}
		return u
	}
	var find func(x int, u *uf) int
	find = func(x int, u *uf) int {
		if u.parent[x] != x {
			u.parent[x] = find(u.parent[x], u)
		}
		return u.parent[x]
	}
	union := func(a, b int, u *uf) {
		pa, pb := find(a, u), find(b, u)
		if pa == pb {
			return
		}
		u.conn -= 1
		if pa > pb {
			pa, pb = pb, pa
		}
		u.parent[pb] = pa
	}
	check := func(x int) bool {
		u := newUf(n)
		for _, ed := range edges {
			if ed[2] > x {
				continue
			}
			union(ed[0], ed[1], u)
		}
		if u.conn <= k {
			return true
		}
		return false
	}
	for l < r {
		mid := l + (r-l)/2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
