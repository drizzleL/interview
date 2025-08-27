package main

func minTime9(n int, edges [][]int, k int) int {
	check := func(x int) bool {
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
		for _, ed := range edges {
			if ed[2] <= x {
				continue
			}
			union(ed[0], ed[1])
		}
		dict := map[int]bool{}
		for i := range parent {
			dict[find(i)] = true
		}
		return len(dict) >= k
	}
	var l, r int
	for _, ed := range edges {
		r = max(r, ed[2])
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
