package main

func regionsBySlashes(grid []string) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	parent := make([]int, m*n*4)
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
	unionHelper := func(i, j, pos int, i2, j2, pos2 int) {
		if i < 0 || j < 0 || i2 < 0 || j2 < 0 || i >= m || i2 >= m || j >= n || j2 >= n {
			return
		}
		a := (i*n+j)*4 + pos
		b := (i2*n+j2)*4 + pos2
		union(a, b)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case ' ':
				unionHelper(i, j, 0, i, j, 1)
				unionHelper(i, j, 0, i, j, 2)
				unionHelper(i, j, 0, i, j, 3)
			case '/':
				unionHelper(i, j, 0, i, j, 3)
				unionHelper(i, j, 1, i, j, 2)
			case '\\':
				unionHelper(i, j, 0, i, j, 1)
				unionHelper(i, j, 2, i, j, 3)
			}
			unionHelper(i, j, 1, i, j+1, 3)
			unionHelper(i, j, 2, i+1, j, 0)
		}
	}
	group := map[int]bool{}
	for i := range parent {
		group[find(i)] = true
	}
	return len(group)
}
