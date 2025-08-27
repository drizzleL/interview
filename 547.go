package main

func findCircleNum(isConnected [][]int) int {
	parent := make([]int, len(isConnected))
	for i := range parent {
		parent[i] = i
	}
	g := len(isConnected)
	var find func(int) int
	find = func(i int) int {
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return parent[i]
	}
	union := func(a, b int) {
		pa, pb := find(a), find(b)
		if pa == pb {
			return
		}
		if pa > pb {
			pa, pb = pb, pa
		}
		g -= 1
		parent[pb] = pa
	}
	for i := 0; i < len(isConnected); i++ {
		for j := i + 1; j < len(isConnected); j++ {
			if isConnected[i][j] == 0 {
				continue
			}
			union(i, j)
		}
	}
	return g
}
