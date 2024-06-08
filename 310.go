package main

func findMinHeightTrees(n int, edges [][]int) []int {
	out := make([]int, n)
	dict := map[int][]int{}
	for _, edge := range edges {
		out[edge[0]] += 1
		out[edge[1]] += 1
		dict[edge[0]] = append(dict[edge[0]], edge[1])
		dict[edge[1]] = append(dict[edge[1]], edge[0])
	}
	var used int
	nodes := []int{}
	for i := 0; i < n; i++ {
		if out[i] == 1 {
			nodes = append(nodes, i)
		}
	}
	seen := make([]bool, n)
	for len(nodes) != 0 && len(nodes) != n-used {
		var next []int
		for _, n := range nodes {
			used += 1
			seen[n] = true
			for _, child := range dict[n] {
				out[child] -= 1
				if out[child] == 1 {
					next = append(next, child)
				}
			}
		}
		nodes = next
	}
	var ret []int
	for i := range seen {
		if !seen[i] {
			ret = append(ret, i)
		}
	}
	return ret
}
