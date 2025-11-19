package main

import "sort"

func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	tmp := make([]int, 0, n)
	helper := func(i, j int) {
		for i2, j2 := i, j; i2 < m && j2 < n; i2, j2 = i2+1, j2+1 {
			tmp = append(tmp, mat[i2][j2])
		}
		sort.Ints(tmp)
		for idx, i2, j2 := 0, i, j; i2 < m && j2 < n; idx, i2, j2 = 0, i2+1, j2+1 {
			mat[i2][j2] = tmp[idx]
		}
		tmp = tmp[:0]
	}
	for j := 0; j < n; j++ {
		helper(0, j)
	}
	for i := 1; i < m; i++ {
		helper(i, 0)
	}
	return mat
}
