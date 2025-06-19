package main

func assignEdgeWeights(edges [][]int) int {
	n := len(edges) + 1
	dict := make([][]int, n)
	for _, ed := range edges {
		dict[ed[0]-1] = append(dict[ed[0]], ed[1]-1)
		dict[ed[1]-1] = append(dict[ed[1]], ed[0]-1)
	}
	depth := 1
	nodes := []int{0}
	seen := make([]bool, n)
	for len(nodes) > 0 {
		var next []int
		for _, node := range nodes {
			seen[node] = true
			for _, child := range dict[node] {
				if seen[child] {
					continue
				}
				seen[child] = true
				next = append(next, child)
			}
		}
		depth += 1
		nodes = next
	}
	return fastPow(2, depth-1, 1)
}
