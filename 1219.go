package main

func getMaximumGold(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	var traverse func(i, j int) int
	traverse = func(i, j int) int {
		if i < 0 || j < 0 || i >= m || j >= n {
			return 0
		}
		if grid[i][j] == 0 || seen[i][j] {
			return 0
		}
		var ret int
		seen[i][j] = true
		ret = max(ret, traverse(i+1, j))
		ret = max(ret, traverse(i, j+1))
		ret = max(ret, traverse(i-1, j))
		ret = max(ret, traverse(i, j-1))
		seen[i][j] = false
		return ret + grid[i][j]
	}
	var ret int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ret = max(ret, traverse(i, j))
		}
	}
	return ret
}
