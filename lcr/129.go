package main

func wordPuzzle(grid [][]byte, target string) bool {
	if len(grid) == 0 {
		return false
	}
	var starts [][2]int
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == target[0] {
				starts = append(starts, [2]int{i, j})
			}
		}
	}
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var traverse func(i, j int, vis []bool, k int) bool
	traverse = func(i, j int, vis []bool, k int) bool {
		if k+1 == len(target) {
			return true
		}
		vis[i*n+j] = true
		for _, dir := range dirs {
			i2, j2 := i+dir[0], j+dir[1]
			if i2 < 0 || j2 < 0 || i2 >= m || j2 >= n {
				continue
			}
			if vis[i2*n+j2] {
				continue
			}
			if grid[i2][j2] != target[k+1] {
				continue
			}
			if traverse(i2, j2, vis, k+1) {
				return true
			}
		}
		vis[i*n+j] = false
		return false
	}

	for _, s := range starts {
		if traverse(s[0], s[1], make([]bool, m*n), 0) {
			return true
		}
	}
	return false
}
