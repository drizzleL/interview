package main

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	parent := make([]int, n)
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
	union := func(i, j int) {
		if i == j {
			return
		}
		pi, pj := find(i), find(j)
		if pi > pj {
			pi, pj = pj, pi
		}
		parent[pj] = pi
	}
	for _, conn := range connections {
		union(conn[0], conn[1])
	}
	seen := make([]bool, n)
	var groups int
	for i := range parent {
		idx := find(i)
		if seen[idx] {
			continue
		}
		seen[idx] = true
		groups += 1
	}
	return groups - 1
}
