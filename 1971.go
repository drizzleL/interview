package main

func validPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}
	seen := make([]bool, n)
	dict := map[int][]int{}
	for _, edge := range edges {
		dict[edge[0]] = append(dict[edge[0]], edge[1])
		dict[edge[1]] = append(dict[edge[1]], edge[0])
	}
	nodes := []int{source}
	seen[source] = true
	for len(nodes) > 0 {
		var next []int
		for _, n := range nodes {
			for _, child := range dict[n] {
				if seen[child] {
					continue
				}
				if child == destination {
					return true
				}
				seen[child] = true
				next = append(next, child)
			}
		}
		nodes = next
	}
	return false
}
