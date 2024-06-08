package main

func numMagicSquaresInside(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	check := func(i, j int) bool {
		if grid[i][j] != 5 {
			return false
		}
		if grid[i-1][j-1]+grid[i+1][j+1] != 10 {
			return false
		}
		if grid[i-1][j+1]+grid[i+1][j-1] != 10 {
			return false
		}
		var flag int
		for k := i - 1; k <= i+1; k++ {
			var sum int
			for l := j - 1; l <= j+1; l++ {
				if grid[k][l] >= 10 {
					return false
				}
				if flag&(1<<grid[k][l]) != 0 {
					return false
				}
				flag |= 1 << grid[k][l]
				sum += grid[k][l]
			}
			if sum != 15 {
				return false
			}
		}
		for l := j - 1; l <= j+1; l++ {
			var sum int
			for k := i - 1; k <= i+1; k++ {
				sum += grid[k][l]
			}
			if sum != 15 {
				return false
			}
		}
		return true
	}
	var ret int
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if check(i, j) {
				ret += 1
			}
		}
	}
	return ret
}
