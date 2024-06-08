package main

func isBipartite(graph [][]int) bool {
	colors := make([]int, len(graph))
	var dfs func(i int, color int) bool
	dfs = func(i int, color int) bool {
		colors[i] = color
		for _, next := range graph[i] {
			if colors[next] == color {
				return false
			}
			if colors[next] == -color {
				continue
			}
			if dfs(next, -color) {
				continue
			}
			return false
		}
		return true
	}
	for i := range graph {
		if colors[i] != 0 {
			continue
		}
		if dfs(i, 1) {
			continue
		}
		return false
	}
	return true
}
