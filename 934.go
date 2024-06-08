package main

func shortestBridge(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	n := len(grid)
	// find first node
	i, j := func() (int, int) {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == 1 {
					return i, j
				}
			}
		}
		return -1, -1
	}()
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	nodes := [][2]int{}
	var pollute func(i, j int)
	pollute = func(i, j int) {
		for _, dir := range [][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
			i2, j2 := i+dir[0], j+dir[1]
			if i2 < 0 || j2 < 0 || i2 >= n || j2 >= n {
				continue
			}
			if visited[i2][j2] {
				continue
			}
			if grid[i2][j2] == 0 {
				continue
			}
			visited[i2][j2] = true
			nodes = append(nodes, [2]int{i2, j2})
			pollute(i2, j2)
		}
	}
	visited[i][j] = true
	nodes = append(nodes, [2]int{i, j})
	pollute(i, j)
	for step := 0; ; step++ {
		var next [][2]int
		for _, node := range nodes {
			i, j := node[0], node[1]
			for _, dir := range [][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
				i2, j2 := i+dir[0], j+dir[1]
				if i2 < 0 || j2 < 0 || i2 >= n || j2 >= n {
					continue
				}
				if visited[i2][j2] {
					continue
				}
				if grid[i2][j2] == 1 { // found
					return step
				}
				visited[i2][j2] = true
				next = append(next, [2]int{i2, j2})
			}
		}
		nodes = next
	}
	return 0
}
