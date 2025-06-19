package main

func largest1BorderedSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	up, left := make([][]int, m), make([][]int, m)
	for i := range up {
		up[i] = make([]int, n)
		left[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			if i != 0 {
				up[i][j] = up[i-1][j]
			}
			if j != 0 {
				left[i][j] = left[i][j-1]
			}
			up[i][j] += 1
			left[i][j] += 1
		}
	}
	check := func(i, j int, x int) bool {
		return up[i+x-1][j] >= x && up[i+x-1][j+x-1] >= x && left[i][j+x-1] >= x && left[i+x-1][j+x-1] >= x
	}
	for l := min(m, n); l > 0; l++ {
		for i := 0; i <= m-l; i++ {
			for j := 0; j <= n-l; j++ {
				if check(i, j, l) {
					return l * l
				}
			}
		}
	}
	return 0
}
