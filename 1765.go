package main

func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	ret := make([][]int, m)
	for i := range ret {
		ret[i] = make([]int, n)
		for j := range ret[i] {
			ret[i][j] = -1
		}
	}
	var nodes [][2]int
	for i := range isWater {
		for j := range isWater[i] {
			if isWater[i][j] == 1 {
				ret[i][j] = 0
				nodes = append(nodes, [2]int{i, j})
			}
		}
	}
	for k := 1; len(nodes) > 0; k++ {
		var next [][2]int
		for _, node := range nodes {
			i, j := node[0], node[1]
			for _, dir := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				i2, j2 := i+dir[0], j+dir[1]
				if i2 < 0 || j2 < 0 || i2 >= m || j2 >= n {
					continue
				}
				if ret[i2][j2] != -1 {
					continue
				}
				ret[i2][j2] = k
				next = append(next, [2]int{i2, j2})
			}
		}
		nodes = next
	}
	return ret
}
