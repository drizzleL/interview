package main

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}
	for j := 0; j < n; j++ {
		for i := 1; i < m; i++ {
			matrix[i][j] += matrix[i-1][j]
		}
	}
	query := func(x1, x2, y2 int) int {
		ret := matrix[x2][y2]
		if x1 > 0 {
			ret -= matrix[x1-1][y2]
		}
		return ret
	}
	var ret int
	for i1 := 0; i1 < m; i1++ {
		for i2 := i1; i2 < m; i2++ {
			dict := map[int]int{}
			dict[0] = 1
			for j := 0; j < n; j++ {
				sum := query(i1, i2, j)
				if v, ok := dict[target-sum]; ok {
					ret += v
				}
				dict[sum] += 1
			}
		}
	}
	return ret
}
