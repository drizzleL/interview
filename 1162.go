package main

func maxDistancev2(grid [][]int) int {
	var nodes [][2]int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				nodes = append(nodes, [2]int{i, j})
				grid[i][j] = -1
			}
		}
	}
	for v := 1; len(nodes) != 0; v++ {
		var next [][2]int
		for _, n := range nodes {
			for _, dir := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				i2, j2 := n[0]+dir[0], n[1]+dir[1]
				if i2 < 0 || j2 < 0 || i2 >= len(grid) || j2 >= len(grid[0]) {
					continue
				}
				if grid[i2][j2] != 0 {
					continue
				}
				grid[i2][j2] = v
				next = append(next, [2]int{i2, j2})
			}
		}
		nodes = next
	}
	var ret int
	for i := range grid {
		for j := range grid[i] {
			ret = max(ret, grid[i][j])
		}
	}
	if ret == 0 {
		return -1
	}
	return ret
}
