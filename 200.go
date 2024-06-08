package main

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	parent := make([]int, m*n)
	for i := range parent {
		parent[i] = i
	}
	findKey := func(i, j int) int {
		return i*n + j
	}
	toCoor := func(x int) (int, int) {
		return x / n, x % n
	}
	var find func(a int) int
	find = func(a int) int {
		if parent[a] != a {
			parent[a] = find(parent[a])
		}
		return parent[a]
	}
	union := func(a, b int) {
		if a == b {
			return
		}
		fa, fb := find(a), find(b)
		if fa > fb {
			fa, fb = fb, fa
		}
		parent[fb] = fa
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			if i != 0 && grid[i-1][j] == '1' {
				union(findKey(i, j), findKey(i-1, j))
			}
			if j != 0 && grid[i][j-1] == '1' {
				union(findKey(i, j), findKey(i, j-1))
			}
		}
	}
	dict := map[int]bool{}
	for i := range parent {
		dict[find(i)] = true
	}
	var ret int
	for k := range dict {
		i, j := toCoor(k)
		if grid[i][j] != '1' {
			continue
		}
		ret += 1
	}
	return ret
}
