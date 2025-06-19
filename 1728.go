package main

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	m, n := len(grid), len(grid[0])
	var originalMouse, originalCat, food [2]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'M' {
				originalMouse = [2]int{i, j}
			}
			if grid[i][j] == 'C' {
				originalCat = [2]int{i, j}
			}
			if grid[i][j] == 'F' {
				food = [2]int{i, j}
			}
		}
	}
	dp := make([][][][][]int, m)
	degree := make([][][][][]int, m)
	for i := range dp {
		dp[i] = make([][][][]int, n)
		degree[i] = make([][][][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([][][]int, m)
			degree[i][j] = make([][][]int, m)
			for k := range dp[i][j] {
				dp[i][j][k] = make([][]int, n)
				degree[i][j][k] = make([][]int, n)
				for l := range dp[i][j][k] {
					dp[i][j][k][l] = make([]int, 2)
					degree[i][j][k][l] = make([]int, 2)
				}
			}
		}
	}
	var q [][5]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '#' {
				continue
			}
			// mouse wins
			if grid[i][j] == 'F' {
				for k := 0; k < m; k++ {
					for l := 0; l < n; l++ {
						if grid[k][l] == '#' || grid[k][l] == 'F' {
							continue
						}
						dp[i][j][k][l][0] = 1
						q = append(q, [5]int{i, j, k, l, 0})
						dp[i][j][k][l][1] = 1
						q = append(q, [5]int{i, j, k, l, 1})
					}
				}
				continue
			}
			// cat wins
			dp[i][j][i][j][0] = 2
			q = append(q, [5]int{i, j, i, j, 0})
			dp[i][j][i][j][1] = 2
			q = append(q, [5]int{i, j, i, j, 1})
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < m; k++ {
				for l := 0; l < n; l++ {
					for _, dir := range [][4]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
						degree[i][j][k][l][0] += 1
						degree[i][j][k][l][1] += 1
						for step := 1; step <= mouseJump; step++ {
							i2, j2 := i+step*dir[0], j+step*dir[1]
							if i2 < 0 || j2 < 0 || i2 >= m || j2 >= n {
								break
							}
							if grid[i2][j2] == '#' {
								break
							}
							degree[i][j][k][l][0] += 1
						}
						for step := 1; step <= catJump; step++ {
							k2, l2 := k+step*dir[0], l+step*dir[1]
							if k2 < 0 || l2 < 0 || k2 >= m || l2 >= n {
								break
							}
							if grid[k2][l2] == '#' || grid[k2][l2] == 'F' {
								break
							}
							degree[i][j][k][l][1] += 1
						}
					}
				}
			}
		}
	}
	for len(q) > 0 {
		top := q[len(q)-1]
		q = q[:len(q)-1]
		mouse, cat, turn := [2]int{top[0], top[1]}, [2]int{top[2], top[3]}, top[4]
		c := dp[mouse[0]][mouse[1]][cat[0]][cat[1]][0]
		if mouse == originalMouse && cat == originalCat && turn == 0 {
			return c == 1
		}
		prevTurn := 1 - turn
		prevMouse, prevCat := mouse, cat
		if dp[mouse[0]][mouse[1]][cat[0]][cat[1]][prevTurn] == 0 { // stays
			q = append(q, [5]int{mouse[0], mouse[1], cat[0], cat[1], prevTurn})
		}
		jump := mouseJump
		obj := mouse
		if prevTurn == 1 {
			jump = catJump
			obj = cat
		}
		for _, dir := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			for i := 1; i <= jump; i++ {
				prevObj := [2]int{obj[0] + dir[0]*i, obj[1] + dir[1]*i}
				if prevObj[0] < 0 || prevObj[1] < 0 || prevObj[0] >= m || prevObj[1] >= n {
					break
				}
				if grid[prevObj[0]][prevObj[1]] == '#' {
					break
				}
				if prevTurn == 0 {
					prevMouse = prevObj
				} else {
					prevCat = prevObj
				}
				if prevCat == food {
					continue
				}
				if dp[prevMouse[0]][prevMouse[1]][cat[0]][cat[1]][prevTurn] != 0 {
					continue
				}
				degree[prevMouse[0]][prevMouse[1]][cat[0]][cat[1]][prevTurn] -= 1
				if prevTurn == 0 && c == 1 || prevTurn == 1 && c == 2 || degree[prevMouse[0]][prevMouse[1]][cat[0]][cat[1]][prevTurn] == 0 {
					q = append(q, [5]int{prevMouse[0], prevMouse[1], prevCat[0], prevCat[1], prevTurn})
					dp[prevMouse[0]][prevMouse[1]][prevCat[0]][prevCat[1]][prevTurn] = c
				}
			}
		}
	}
	return false
}
