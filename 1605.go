package main

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	var i, j int
	ret := make([][]int, len(rowSum))
	for i := range ret {
		ret[i] = make([]int, len(colSum))
	}
	for i < len(rowSum) && j < len(colSum) {
		if rowSum[i] <= colSum[j] {
			ret[i][j] = rowSum[i]
			colSum[j] -= rowSum[i]
			i += 1
		} else {
			ret[i][j] = colSum[j]
			rowSum[i] -= colSum[j]
			j += 1
		}
	}
	return ret
}
