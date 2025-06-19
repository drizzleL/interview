package main

func numberOfComponents(properties [][]int, k int) int {
	parent := make([]int, len(properties))
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
	dict := make([]map[int]bool, len(properties))
	for i := range properties {
		dict[i] = map[int]bool{}
		for _, v := range properties[i] {
			dict[i][v] = true
		}
	}
	check := func(i, j int) bool {
		var cnt int
		for v := range dict[i] {
			if dict[j][v] {
				cnt += 1
			}
		}
		return cnt >= k
	}
	for i := 0; i < len(properties); i++ {
		for j := i + 1; j < len(properties); j++ {
			if check(i, j) {
				union(i, j)
			}
		}
	}
	groups := map[int]bool{}
	for i := range parent {
		groups[find(i)] = true
	}
	return len(groups)
}
