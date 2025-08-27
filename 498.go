package main

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	ret := make([]int, 0, m*n)
	for i, j := 0, 0; i < m && j < n; {
		for i < m && j < n {
			ret = append(ret, mat[i][j])
			if j == n-1 {
				i += 1
				break
			}
			if i == 0 {
				j += 1
				break
			}
			i -= 1
			j += 1
		}
		for i < m && j < n {
			ret = append(ret, mat[i][j])
			if i == m-1 {
				j += 1
				break
			}
			if j == 0 {
				i += 1
				break
			}
			i += 1
			j -= 1
		}
	}
	return ret
}
