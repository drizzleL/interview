package main

func minimumJumps(forbidden []int, a int, b int, x int) int {
	seen := make(map[[2]int]bool)
	nodes := [][2]int{{0, 0}}
	maxVal := x
	for _, forb := range forbidden {
		seen[[2]int{forb, 0}] = true
		seen[[2]int{forb, 1}] = true
		maxVal = max(maxVal, forb)
	}
	for step := 0; len(nodes) > 0; step++ {
		var next [][2]int
		for _, node := range nodes {
			if node[0] == x {
				return step
			}
			forward := node[0] + a
			if forward <= maxVal && !seen[[2]int{forward, 0}] {
				seen[[2]int{forward, 0}] = true
				next = append(next, [2]int{forward, 0})
			}
			backward := node[0] - b
			if node[1] == 0 && backward >= 0 && !seen[[2]int{backward, 1}] {
				seen[[2]int{backward, 1}] = true
				next = append(next, [2]int{backward, 1})
			}
		}
		nodes = next
	}
	return -1
}
