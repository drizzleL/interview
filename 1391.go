package main

func hasValidPath(grid [][]int) bool {
	if len(grid) == 1 && len(grid[0]) == 1 {
		return true
	}
	dirMapping := []int{2, 3, 0, 1}
	endMapping := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	dirs := map[int]map[int]int{
		1: {
			1: 3,
			3: 1,
		},
		2: {
			0: 2,
			2: 0,
		},
		3: {
			2: 3,
			3: 2,
		},
		4: {
			1: 2,
			2: 1,
		},
		5: {
			0: 3,
			3: 0,
		},
		6: {
			0: 1,
			1: 0,
		},
	}
	var helper func(i, j, from int) bool
	helper = func(i, j, from int) bool {
		if i == 0 && j == 0 {
			return false
		}
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
			return false
		}
		d := grid[i][j]
		dir := dirs[d]
		end, ok := dir[from]
		if !ok {
			return false
		}
		if i == len(grid)-1 && j == len(grid[0])-1 {
			return true
		}
		newFrom := dirMapping[end]
		newI, newJ := i+endMapping[end][0], j+endMapping[end][1]
		return helper(newI, newJ, newFrom)
	}
	for _, end := range dirs[grid[0][0]] {
		newFrom := dirMapping[end]
		newI, newJ := endMapping[end][0], endMapping[end][1]
		if helper(newI, newJ, newFrom) {
			return true
		}
	}
	return false
}
