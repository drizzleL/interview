package main

func minimumArea(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var flag bool
	l, r, t, b := n, 0, m, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			flag = true
			l = min(l, j)
			r = max(r, j)
			t = min(t, i)
			b = max(b, i)
		}
	}
	if !flag {
		return 0
	}
	return (r - l + 1) * (b - t + 1)
}
