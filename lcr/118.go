package main

func findRedundantConnection(edges [][]int) []int {
	dict := map[int][]int{}
	degree := make([]int, len(edges)+1)
	for _, e := range edges {
		i, j := e[0], e[1]
		dict[i] = append(dict[i], j)
		dict[j] = append(dict[j], i)
		degree[i]++
		degree[j]++
	}
	var nodes []int
	for i, v := range degree {
		if v == 1 {
			nodes = append(nodes, i)
		}
	}
	for len(nodes) != 0 {
		var next []int
		for _, n := range nodes {
			degree[n]--
			for _, child := range dict[n] {
				if degree[child] == 0 {
					continue
				}
				degree[child]--
				if degree[child] == 1 {
					next = append(next, child)
				}
			}
		}
		nodes = next
	}
	cycleDict := make([]bool, len(edges)+1)
	for i, v := range degree {
		if v == 0 {
			continue
		}
		cycleDict[i] = true
	}
	for i := len(edges) - 1; i >= 0; i-- {
		e := edges[i]
		if cycleDict[e[0]] && cycleDict[e[1]] {
			return e
		}
	}
	return nil
}
