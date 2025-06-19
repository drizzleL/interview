package main

func countPathsWithXorValue(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	valGrid := make([][]map[int]int, m)
	for i := range valGrid {
		valGrid[i] = make([]map[int]int, n)
		for j := range valGrid[i] {
			valGrid[i][j] = map[int]int{}
		}
	}
	valGrid[0][0] = map[int]int{grid[0][0]: 1}
	for j := 1; j < n; j++ {
		for k, v := range valGrid[0][j-1] {
			valGrid[0][j][k^grid[0][j]] += v
		}
	}
	for i := 1; i < m; i++ {
		for k, v := range valGrid[i-1][0] {
			valGrid[i][0][k^grid[i][0]] += v
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			for k, v := range valGrid[i][j-1] {
				valGrid[i][j][k^grid[i][j]] += v
				valGrid[i][j][k^grid[i][j]] %= 1e9 + 7
			}
			for k, v := range valGrid[i-1][j] {
				valGrid[i][j][k^grid[i][j]] += v
				valGrid[i][j][k^grid[i][j]] %= 1e9 + 7
			}
		}
	}
	return valGrid[m-1][n-1][k]
}
