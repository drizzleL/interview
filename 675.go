package main

import "sort"

func cutOffTree(forest [][]int) int {
	if forest[0][0] == 0 {
		return -1
	}
	m, n := len(forest), len(forest[0])
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	var zero int
	var nodes [][3]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if forest[i][j] == 0 {
				zero += 1
				continue
			}
			if forest[i][j] == 1 {
				continue
			}
			nodes = append(nodes, [3]int{forest[i][j], i, j})
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		seen[i][j] = true
		ret := 1
		for _, dir := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			i2, j2 := i+dir[0], j+dir[1]
			if i2 < 0 || j2 < 0 || i2 >= m || j2 >= n {
				continue
			}
			if seen[i2][j2] {
				continue
			}
			seen[i2][j2] = true
			if forest[i2][j2] == 0 {
				continue
			}
			ret += dfs(i2, j2)
		}
		return ret
	}
	canWalk := dfs(0, 0)
	if canWalk+zero != m*n {
		return -1
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i][0] < nodes[j][0]
	})
	walk := func(lastX, lastY, x, y int) int {
		seen := make([][]bool, m)
		for i := range seen {
			seen[i] = make([]bool, n)
		}
		points := [][2]int{{lastX, lastY}}
		for step := 1; len(points) > 0; step++ {
			var next [][2]int
			for _, p := range points {
				for _, dir := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
					i2, j2 := p[0]+dir[0], p[1]+dir[1]
					if i2 < 0 || j2 < 0 || i2 >= m || j2 >= n {
						continue
					}
					if seen[i2][j2] {
						continue
					}
					seen[i2][j2] = true
					if forest[i2][j2] == 0 {
						continue
					}
					if i2 == x && j2 == y {
						return step
					}
					next = append(next, [2]int{i2, j2})
				}
			}
			points = next
		}
		return 0
	}
	var ret int
	var lastX, lastY int
	if nodes[0][1] == 0 && nodes[0][2] == 0 { // cut before step
		nodes = nodes[1:]
	}
	for _, node := range nodes {
		x, y := node[1], node[2]
		ret += walk(lastX, lastY, x, y)
		lastX, lastY = x, y
	}
	return ret
}
