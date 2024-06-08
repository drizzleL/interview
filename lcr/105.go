package main

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	parent := make([]int, m*n)
	for i := range parent {
		parent[i] = i
	}
	var find func(key int) int
	find = func(key int) int {
		if parent[key] != key {
			parent[key] = find(parent[key])
		}
		return parent[key]
	}
	union := func(k1, k2 int) {
		if k1 > k2 {
			k1, k2 = k2, k1
		}
		parent[find(k2)] = find(k1)
	}
	getKey := func(i, j int) int {
		return i*n + j
	}
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	check := func(i, j int) {
		for _, dir := range dirs {
			i2, j2 := i+dir[0], j+dir[1]
			if i2 < 0 || i2 >= m || j2 < 0 || j2 >= n {
				continue
			}
			if grid[i2][j2] == 0 {
				continue
			}
			union(getKey(i, j), getKey(i2, j2))
		}
	}
	var flag bool
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			flag = true
			check(i, j)
		}
	}
	if !flag {
		return 0
	}
	dict := map[int]int{}
	var ret int
	for _, p := range parent {
		dict[find(p)]++
		if dict[find(p)] > ret {
			ret = dict[find(p)]
		}
	}
	return ret
}
