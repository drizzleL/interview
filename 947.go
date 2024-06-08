package main

func removeStones(stones [][]int) int {
	parent := make([]int, len(stones))
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
	rows, cols := map[int][]int{}, map[int][]int{}
	for i, st := range stones {
		rows[st[0]] = append(rows[st[0]], i)
		cols[st[1]] = append(cols[st[1]], i)
	}
	for _, rowGroup := range rows {
		for i := 1; i < len(rowGroup); i++ {
			union(rowGroup[i], rowGroup[i-1])
		}
	}
	for _, group := range cols {
		for i := 1; i < len(group); i++ {
			union(group[i], group[i-1])
		}
	}
	group := map[int]bool{}
	for _, v := range parent {
		group[find(v)] = true
	}
	return len(stones) - len(group)
}
