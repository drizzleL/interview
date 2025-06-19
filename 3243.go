package main

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	dict := make([][]int, n)
	for i := range dict {
		dict[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			dict[i][j] = j - i
		}
	}
	ret := make([]int, len(queries))
	for k, q := range queries {
		dict[q[0]][q[1]] = 1
		for i := 0; i <= q[0]; i++ {
			for j := q[1]; j < n; j++ {
				dict[i][j] = min(dict[i][j], dict[i][q[0]]+1+dict[q[1]][j])
			}
		}
		ret[k] = dict[0][n-1]
	}
	return ret
}
