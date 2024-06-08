package main

func minReorder(n int, connections [][]int) int {
	outDict := map[int][]int{}
	inDict := map[int][]int{}
	for _, c := range connections {
		u, v := c[0], c[1]
		outDict[u] = append(outDict[u], v)
		inDict[v] = append(inDict[v], u)
	}
	seen := make([]bool, n)
	nodes := []int{0}
	seen[0] = true
	var ret int
	for len(nodes) != 0 {
		var next []int
		for _, node := range nodes {
			for _, child := range outDict[node] {
				if seen[child] {
					continue
				}
				ret += 1
				next = append(next, child)
				seen[child] = true
			}
			for _, child := range inDict[node] {
				if seen[child] {
					continue
				}
				next = append(next, child)
				seen[child] = true
			}
		}
		nodes = next
	}
	return ret
}
