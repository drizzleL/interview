package main

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	cache := map[int]int{}
	var helper func(i, j int) int
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	helper = func(i, j int) (ret int) {
		if c, ok := cache[i*n+j]; ok {
			return c
		}
		defer func() {
			cache[i*n+j] = ret
		}()
		for _, dir := range dirs {
			i2, j2 := i+dir[0], j+dir[1]
			if i2 < 0 || j2 < 0 || i2 >= m || j2 >= n {
				continue
			}
			if matrix[i2][j2] <= matrix[i][j] {
				continue
			}
			ret = max(ret, helper(i2, j2))
		}
		ret += 1
		return ret
	}
	var ret int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ret = max(ret, helper(i, j))
		}
	}
	return ret
}
