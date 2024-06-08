package main

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	if start == target {
		return true
	}
	dict := make([][]int, n)
	for _, c := range graph {
		dict[c[0]] = append(dict[c[0]], c[1])
	}
	vis := make([]bool, n)
	vis[start] = true
	nodes := []int{start}
	for len(nodes) != 0 {
		var next []int
		for _, node := range nodes {
			for _, e := range dict[node] {
				if e == target {
					return true
				}
				if vis[e] {
					continue
				}
				vis[e] = true
				next = append(next, e)
			}
		}
		nodes = next
	}
	return false
}
