package main

func rangeAddQueries(n int, queries [][]int) [][]int {
	ret := make([][]int, n)
	for i := range ret {
		ret[i] = make([]int, n)
	}
	for _, q := range queries {
		r1, c1, r2, c2 := q[0], q[1], q[2], q[3]
		ret[r1][c1] += 1
		if c2+1 < n {
			ret[r1][c2+1] -= 1
		}
		if r2+1 < n {
			ret[r2+1][c1] -= 1
		}
		if r2+1 < n && c2+1 < n {
			ret[r2+1][c2+1] += 1
		}
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			ret[i][j] += ret[i-1][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			ret[i][j] += ret[i][j-1]
		}
	}
	return ret
}
