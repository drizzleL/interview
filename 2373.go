package main

func largestLocal(grid [][]int) [][]int {
	size := len(grid)
	ret := make([][]int, size-2)
	for i := range ret {
		ret[i] = make([]int, size-2)
	}
	for i := 1; i < size-1; i++ {
		for j := 1; j < size-1; j++ {
			ret[i-1][j-1] = max(ret[i-1][j-1], grid[i-1][j-1])
			ret[i-1][j-1] = max(ret[i-1][j-1], grid[i-1][j])
			ret[i-1][j-1] = max(ret[i-1][j-1], grid[i][j-1])
			ret[i-1][j-1] = max(ret[i-1][j-1], grid[i][j])
			ret[i-1][j-1] = max(ret[i-1][j-1], grid[i+1][j+1])
			ret[i-1][j-1] = max(ret[i-1][j-1], grid[i][j+1])
			ret[i-1][j-1] = max(ret[i-1][j-1], grid[i+1][j])
			ret[i-1][j-1] = max(ret[i-1][j-1], grid[i+1][j-1])
			ret[i-1][j-1] = max(ret[i-1][j-1], grid[i-1][j+1])
		}
	}
	return ret
}
