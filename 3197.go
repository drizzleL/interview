package main

func minimumSum3(grid [][]int) int {
	single := func(grid [][]int, upFrom int, leftFrom int) int {
		up, down, left, right := len(grid)-upFrom, -1, len(grid[0])-leftFrom, -1
		for i := upFrom; i < len(grid); i++ {
			for j := leftFrom; j < len(grid[0]); j++ {
				if grid[i][j] == 0 {
					continue
				}
				up = min(up, i)
				down = max(down, i)
				left = min(left, j)
				right = max(right, j)
			}
		}
		if up == len(grid) {
			return 0
		}
		return (down - up + 1) * (right - left + 1)
	}
	split2 := func(grid [][]int) int {
		if len(grid) == 0 {
			return 0
		}
		ret := len(grid) * len(grid[0])
		startI := -1
		left, right := len(grid[0]), -1
		for i := 0; i < len(grid); i++ {
			var found bool
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j] == 0 {
					continue
				}
				found = true
				left = min(left, j)
				right = max(right, j)
			}
			if found && startI == -1 { // init startI
				startI = i
			}
			if found {
				up := (i - startI + 1) * (right - left + 1)
				ret = min(ret, up+single(grid, i+1, 0))
			}
		}
		startJ := -1
		up, down := len(grid), -1
		for j := 0; j < len(grid[0]); j++ {
			var found bool
			for i := 0; i < len(grid); i++ {
				if grid[i][j] == 0 {
					continue
				}
				found = true
				up = min(up, i)
				down = max(down, i)
			}
			if found && startJ == -1 { // init startJ
				startJ = j
			}
			if found {
				left := (j - startJ + 1) * (down - up + 1)
				ret = min(ret, left+single(grid, 0, j+1))
			}
		}
		return ret
	}
	h2 := func(grid [][]int) int {
		ret := len(grid) * len(grid[0])
		startI := -1
		left, right := len(grid[0]), -1
		for i := 0; i < len(grid); i++ {
			var found bool
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j] == 0 {
					continue
				}
				found = true
				left = min(left, j)
				right = max(right, j)
			}
			if found && startI == -1 { // init startI
				startI = i
			}
			if found {
				up := (i - startI + 1) * (right - left + 1)
				ret = min(ret, up+split2(grid[i+1:]))
			}
		}
		return ret
	}
	helper := func(grid [][]int) int {
		ret := h2(grid)
		for i, j := 0, len(grid)-1; i < j; i, j = i+1, j-1 {
			grid[i], grid[j] = grid[j], grid[i]
		}
		ret = min(ret, h2(grid))
		return ret
	}
	g2 := make([][]int, len(grid[0]))
	for i := range g2 {
		g2[i] = make([]int, len(grid))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			g2[j][i] = grid[i][j]
		}
	}
	return min(helper(grid), helper(g2))
}
