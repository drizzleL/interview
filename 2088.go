package main

func countPyramids(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	sums := make([][]int, m)
	for i := range sums {
		sums[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sums[i][j+1] = sums[i][j] + grid[i][j]
		}
	}
	check := func(line int, i, j int) bool {
		if line < 0 || line >= m || i < 0 || j < 0 || i >= n || j >= n {
			return false
		}
		return sums[line][j+1]-sums[line][i] == j-i+1
	}
	var ret int
	lineSum := make([]int, n)
	for i := 0; i < m; i++ {
		newLineSum := make([]int, n)
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 {
				continue
			}
			newLineSum[j] = max(max(lineSum[j-1]-1, lineSum[j+1]-1), 0)
			for check(i+newLineSum[j]+1, j-newLineSum[j]-1, j+newLineSum[j]+1) {
				newLineSum[j] += 1
			}
			ret += newLineSum[j]
		}
		lineSum = newLineSum
	}
	lineSum = make([]int, n)
	for i := m - 1; i >= 0; i-- {
		newLineSum := make([]int, n)
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 {
				continue
			}
			newLineSum[j] = max(max(lineSum[j-1]-1, lineSum[j+1]-1), 0)
			for check(i-newLineSum[j]-1, j-newLineSum[j]-1, j+newLineSum[j]+1) {
				newLineSum[j] += 1
			}
			ret += newLineSum[j]
		}
		lineSum = newLineSum
	}
	return ret
}
