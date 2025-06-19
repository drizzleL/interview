package main

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	for _, w := range walls {
		grid[w[0]][w[1]] = 2
	}
	for _, g := range guards {
		grid[g[0]][g[1]] = 1
	}
	for _, g := range guards {
		i, j := g[0], g[1]
		for a, b := i-1, j; a >= 0 && grid[a][b] <= 0; a-- {
			grid[a][b] = -1
		}
		for a, b := i, j-1; b >= 0 && grid[a][b] <= 0; b-- {
			grid[a][b] = -1
		}
		for a, b := i+1, j; a < m && grid[a][b] <= 0; a++ {
			grid[a][b] = -1
		}
		for a, b := i, j+1; b < n && grid[a][b] <= 0; b++ {
			grid[a][b] = -1
		}

	}
	var ret int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != 0 {
				continue
			}
			ret += 1
		}
	}
	return ret
}
