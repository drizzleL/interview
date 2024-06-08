package main

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	nodes := [][2]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				nodes = append(nodes, [2]int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}
	dirs := [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	getAround := func(i, j int) [][2]int {
		var ret [][2]int
		for _, dir := range dirs {
			i2, j2 := dir[0]+i, dir[1]+j
			if i2 < 0 || j2 < 0 || i2 >= m || j2 >= n {
				continue
			}
			ret = append(ret, [2]int{i2, j2})
		}
		return ret
	}
	val := 1
	for len(nodes) != 0 {
		var next [][2]int
		for _, n := range nodes {
			arounds := getAround(n[0], n[1])
			for _, v := range arounds {
				if mat[v[0]][v[1]] != -1 {
					continue
				}
				mat[v[0]][v[1]] = val
				next = append(next, v)
			}
		}
		nodes = next
		val++
	}
	return mat
}
