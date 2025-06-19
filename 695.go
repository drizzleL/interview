package main

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	parent := make([]int, m*n)
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
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			if i != 0 && grid[i-1][j] == 1 {
				union((i-1)*n+j, i*n+j)
			}
			if j != 0 && grid[i][j-1] == 1 {
				union(i*n+j-1, i*n+j)
			}
		}
	}
	var ret int
	groups := map[int]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			key := find(i*n + j)
			groups[key] += 1
			ret = max(ret, groups[key])
		}
	}
	return ret
}
