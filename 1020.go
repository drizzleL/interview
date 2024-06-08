package main

func numEnclaves(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	if m < 1 || n <= 1 {
		return 0
	}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var helper func(i, j int)
	helper = func(i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		if visited[i][j] {
			return
		}
		visited[i][j] = true
		if grid[i][j] != 1 {
			return
		}
		grid[i][j] = 2
		helper(i-1, j)
		helper(i+1, j)
		helper(i, j+1)
		helper(i, j-1)
	}
	for j := 0; j < n; j++ {
		helper(0, j)
		helper(m-1, j)
	}
	for i := 0; i < m; i++ {
		helper(i, 0)
		helper(i, n-1)
	}
	var ret int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ret++
			}
		}
	}
	return ret
}
