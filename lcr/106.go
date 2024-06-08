package main

func isBipartite(graph [][]int) bool {
	if len(graph) == 0 {
		return false
	}
	var groups int
	dict := make([]int, len(graph))
	for i := range dict {
		dict[i] = -1
	}
	var ans bool
	traverse := func(i int) {
		if dict[i] != -1 {
			return
		}
		if len(graph[i]) == 0 {
			dict[i] = groups
			groups++
			return
		}
		idx := groups
		mirrorIdx := groups + 1
		groups += 2
		nodes := []int{i}
		for len(nodes) != 0 {
			var next []int
			for _, n := range nodes {
				dict[n] = idx
				for _, child := range graph[n] {
					if dict[child] == idx {
						ans = true // early break
						return
					}
					if dict[child] != -1 {
						continue
					}
					next = append(next, child)
				}
			}
			idx, mirrorIdx = mirrorIdx, idx
			nodes = next
		}
	}
	for i := range graph {
		traverse(i)
	}
	return !ans
}
