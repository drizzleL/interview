package main

func numSubmat(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			if mat[i][j] == 0 {
				continue
			}
			mat[i][j] += mat[i][j-1]
		}
	}
	var ret int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				continue
			}
			v := mat[i][j]
			for k := i; v > 0 && k < m; k++ {
				v = min(v, mat[k][j])
				ret += v * (k - i + 1)
			}

		}
	}
	return ret
}
