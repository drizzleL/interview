package main

func minMoves3(matrix []string) int {
	m, n := len(matrix), len(matrix[0])
	if matrix[0][0] == '#' || matrix[m-1][n-1] == '#' {
		return -1
	}
	dict := map[byte][][2]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '.' || matrix[i][j] == '#' {
				continue
			}
			dict[matrix[i][j]] = append(dict[matrix[i][j]], [2]int{i, j})
		}
	}
	seen := make([]bool, m*n)
	var nodes []int
	if matrix[0][0] == '.' {
		seen[0] = true
		nodes = append(nodes, 0)
	} else {
		for _, child := range dict[matrix[0][0]] {
			seen[child[0]*n+child[1]] = true
			nodes = append(nodes, child[0]*n+child[1])
		}
	}
	for step := 0; len(nodes) != 0; step++ {
		var next []int
		for _, node := range nodes {
			x, y := node/n, node%n
			if x == m-1 && y == n-1 {
				return step
			}
			for _, dir := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				x2, y2 := x+dir[0], y+dir[1]
				if x2 < 0 || y2 < 0 || x2 >= m || y2 >= n {
					continue
				}
				if seen[x2*n+y2] {
					continue
				}
				if matrix[x2][y2] == '#' {
					continue
				}
				switch matrix[x2][y2] {
				case '.':
					seen[x2*n+y2] = true
					next = append(next, x2*n+y2)
				default:
					for _, child := range dict[matrix[x2][y2]] {
						seen[child[0]*n+child[1]] = true
						next = append(next, child[0]*n+child[1])
					}
				}
			}
		}
		nodes = next
	}
	return -1
}
