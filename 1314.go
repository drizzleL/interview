package main

func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	helper := func(i, j int) int {
		if i < 0 || j < 0 || i >= m || j >= n {
			return 0
		}
		return mat[i][j]
	}
	getSum := func(l, r, up, down int) int {
		l = max(0, l)
		up = max(0, up)
		r = min(n-1, r)
		down = min(m-1, down)
		return helper(down, r) + helper(up-1, l-1) - helper(up-1, r) - helper(down, l-1)
	}
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			mat[i][j] += mat[i][j-1]
		}
	}
	for j := 0; j < n; j++ {
		for i := 1; i < m; i++ {
			mat[i][j] += mat[i-1][j]
		}
	}
	ret := make([][]int, m)
	for i := range ret {
		ret[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ret[i][j] = getSum(j-k, j+k, i-k, i+k)
		}
	}
	return ret
}
