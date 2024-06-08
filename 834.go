package main

func sumOfDistancesInTree(n int, edges [][]int) []int {
	dict := [][]int{}
	out := make([]int, n)
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], ed[1])
		dict[ed[1]] = append(dict[ed[1]], ed[0])
		out[ed[0]] += 1
		out[ed[1]] += 1
	}
	var nodes []int
	for i := range out {
		if out[i] == 1 {
			nodes = append(nodes, i)
		}
	}
	accu := make([]int, n)
	for len(nodes) != 0 {
		var next []int
		for _, n := range nodes {
			accu[n] += 1
			out[n] -= 1
			for _, child := range dict[n] {
				out[child] -= 1
				if out[child] == 0 {
					next = append(next, child)
				}
			}
		}
		nodes = next
	}
	ret := make([]int, n)
	return ret
}
