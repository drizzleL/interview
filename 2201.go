package main

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, n)
		for j := range mat[i] {
			mat[i][j] = -1
		}
	}
	dict := make([]int, len(artifacts))
	for i, arti := range artifacts {
		for row := arti[0]; row <= arti[2]; row++ {
			for col := arti[1]; col <= arti[3]; col++ {
				mat[row][col] = i
				dict[i] += 1
			}
		}
	}
	var ret int
	for _, d := range dig {
		if mat[d[0]][d[1]] == -1 {
			continue
		}
		idx := mat[d[0]][d[1]]
		dict[idx] -= 1
		if dict[idx] == 0 {
			ret += 1
		}
	}
	return ret
}
