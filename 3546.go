package main

func canPartitionGrid(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	var sum int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum += grid[i][j]
		}
	}
	var sum1, sum2 int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum1 += grid[i][j]
		}
		if sum1*2 == sum {
			return true
		}
	}
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			sum2 += grid[i][j]
		}
		if sum1*2 == sum {
			return true
		}
	}
	return false
}
