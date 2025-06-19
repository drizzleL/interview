package main

func reachableNodes(n int, edges [][]int, restricted []int) int {
	var ret int
	dict := make([][]int, n)
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], ed[1])
		dict[ed[1]] = append(dict[ed[1]], ed[0])
	}
	seen := make([]bool, n)
	for _, res := range restricted {
		seen[res] = true
	}
	nodes := []int{0}
	for len(nodes) > 0 {
		var next []int
		for _, node := range nodes {
			seen[node] = true
			ret += 1
			for _, child := range dict[node] {
				if seen[child] {
					continue
				}
				next = append(next, child)
			}
		}
		nodes = next
	}
	return ret
}
