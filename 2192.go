package main

func getAncestors(n int, edges [][]int) [][]int {
	dict := map[int][]int{}
	in := make([]int, n)
	for _, ed := range edges {
		in[ed[1]] += 1
		dict[ed[0]] = append(dict[ed[0]], ed[1])
	}
	var nodes []int
	for i, v := range in {
		if v == 0 {
			nodes = append(nodes, i)
		}
	}
	retDict := make([][]bool, n)
	for i := range retDict {
		retDict[i] = make([]bool, n)
	}
	for len(nodes) > 0 {
		var next []int
		for _, n := range nodes {
			for _, child := range dict[n] {
				in[child] -= 1
				retDict[child][n] = true
				for j, ok := range retDict[n] {
					if !ok {
						continue
					}
					retDict[child][j] = true
				}
				if in[child] == 0 {
					next = append(next, child)
				}
			}
		}
		nodes = next
	}
	ret := make([][]int, n)
	for i, v := range retDict {
		for j, ok := range v {
			if !ok {
				continue
			}
			ret[i] = append(ret[i], j)
		}
	}
	return ret
}
