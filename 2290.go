package main

func minimumObstacles(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	nodes := [][2]int{{0, 0}}
	seen[0][0] = true
	var ret int
	if grid[0][0] == 1 {
		ret += 1
	}
	if m == 1 && n == 1 {
		return ret
	}
	for ; len(nodes) != 0; ret++ {
		var next [][2]int
		for i := 0; i < len(nodes); i++ {
			node := nodes[i]
			for _, dir := range [][4]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				e1, e2 := node[0]+dir[0], node[1]+dir[1]
				if e1 < 0 || e2 < 0 || e1 >= m || e2 >= n {
					continue
				}
				if seen[e1][e2] {
					continue
				}
				if e1 == m-1 && e2 == n-1 {
					return ret + grid[e1][e2]
				}
				seen[e1][e2] = true
				if grid[e1][e2] == 0 {
					nodes = append(nodes, [2]int{e1, e2})
				} else {
					next = append(next, [2]int{e1, e2})
				}
			}
		}
		nodes = next
	}
	return ret
}
