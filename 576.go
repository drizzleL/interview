package main

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	cache := make([][][]int, m)
	for i := range cache {
		cache[i] = make([][]int, n)
		for j := range cache[i] {
			cache[i][j] = make([]int, maxMove+1)
			for k := range cache[i][j] {
				cache[i][j][k] = -1
			}
		}
	}
	var helper func(i, j int, leftMove int) int
	helper = func(i, j int, leftMove int) (ret int) {
		if i < 0 || j < 0 || i >= m || j >= n {
			return 1
		}
		if leftMove == 0 {
			return 0
		}
		if c := cache[i][j][leftMove]; c != -1 {
			return c
		}
		defer func() {
			cache[i][j][leftMove] = ret
		}()
		ret += helper(i+1, j, leftMove-1)
		ret += helper(i, j+1, leftMove-1)
		ret += helper(i, j-1, leftMove-1)
		ret += helper(i-1, j, leftMove-1)
		return ret % (1e9 + 7)
	}
	return helper(startRow, startColumn, maxMove)
}
