package main

func firstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	rows, cols := make([]int, m), make([]int, n)
	dict := make([][2]int, len(arr)+1)
	for i := range mat {
		for j := range mat[i] {
			dict[mat[i][j]] = [2]int{i, j}
		}
	}
	for i, v := range arr {
		x, y := dict[v][0], dict[v][1]
		rows[x] += 1
		cols[y] += 1
		if rows[x] == m || cols[y] == n {
			return i
		}
	}
	return 0
}
