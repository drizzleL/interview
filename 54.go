package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	var i, j int
	ret := make([]int, 0, m*n)
	ret = append(ret, matrix[i][j])
	up, bottom := 0, m-1
	l, r := 0, n-1
	for len(ret) < m*n {
		for j < r {
			j++
			ret = append(ret, matrix[i][j])
			if j == r {
				up++
				break
			}
		}
		if len(ret) == m*n {
			break
		}
		for i < bottom {
			i++
			ret = append(ret, matrix[i][j])
			if i == bottom {
				r--
				break
			}
		}
		if len(ret) == m*n {
			break
		}
		for j > l {
			j--
			ret = append(ret, matrix[i][j])
			if j == l {
				bottom--
				break
			}
		}
		if len(ret) == m*n {
			break
		}
		for i > up {
			i--
			ret = append(ret, matrix[i][j])
			if i == up {
				l++
				break
			}
		}
	}
	return ret
}
