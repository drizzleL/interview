package main

func maxPathScore(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	check := func(v1, v2 int) int {
		if v1 == -1 {
			return v2
		}
		if v2 == -1 {
			return v1
		}
		return max(v1, v2)
	}
	cache := make([][][]int, m)
	for i := range cache {
		cache[i] = make([][]int, n)
		for j := range cache[i] {
			cache[i][j] = make([]int, k+1)
			for l := range cache[i][j] {
				cache[i][j][l] = -2
			}
		}
	}
	var helper func(i, j int, leftK int) int
	helper = func(i, j int, leftK int) (ret int) {
		if i >= m || j >= n {
			return -1
		}
		if cache[i][j][leftK] != -2 {
			return cache[i][j][leftK]
		}
		defer func() {
			cache[i][j][leftK] = ret
		}()
		ret = grid[i][j]
		if grid[i][j] != 0 && leftK == 0 {
			return -1
		}
		nextK := leftK
		if grid[i][j] != 0 {
			nextK -= 1
		}
		if i == m-1 && j == n-1 {
			return ret
		}
		val := check(helper(i+1, j, nextK), helper(i, j+1, nextK))
		if val == -1 {
			return -1
		}
		ret += val
		return ret
	}
	return helper(0, 0, k)
}
