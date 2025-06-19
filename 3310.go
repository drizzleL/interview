package main

func remainingMethods(n int, k int, invocations [][]int) []int {
	dict := make([][]int, n)
	for _, inv := range invocations {
		dict[inv[0]] = append(dict[inv[0]], inv[1])
	}
	seen := make([]int, n) // 0: default 1: del 2: recover
	nodes := []int{k}
	color := func(nodes []int, color int, find int) bool {
		var flag bool
		for len(nodes) != 0 {
			var next []int
			for _, node := range nodes {
				seen[node] = color
				for _, child := range dict[node] {
					if seen[child] == find {
						flag = true
					}
					if seen[child] == color {
						continue
					}
					seen[child] = color
					next = append(next, child)
				}
			}
			nodes = next
		}
		return flag
	}
	color(nodes, 1, 0)
	var ret []int
	nodes = nodes[:0]
	for i, v := range seen {
		if v == 0 {
			nodes = append(nodes, i)
		}
	}
	find := color(nodes, 2, 1)
	if find { // return all
		for i := 0; i < n; i++ {
			ret = append(ret, i)
		}
		return ret
	}
	for i, v := range seen {
		if v != 1 {
			ret = append(ret, i)
		}
	}
	return ret
}
