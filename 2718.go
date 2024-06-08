package main

func matrixSumQueries(n int, queries [][]int) int64 {
	rows, cols := make([]bool, n), make([]bool, n)
	var rowsCnt, colsCnt int
	var ret int
	for i := len(queries) - 1; i >= 0; i-- {
		q := queries[i]
		switch q[0] {
		case 0:
			if rows[q[1]] {
				continue
			}
			ret += q[2] * (n - colsCnt)
			rows[q[1]] = true
			rowsCnt += 1
		case 1:
			if cols[q[1]] {
				continue
			}
			ret += q[2] * (n - rowsCnt)
			cols[q[1]] = true
			colsCnt += 1
		}
	}
	return int64(ret)
}
