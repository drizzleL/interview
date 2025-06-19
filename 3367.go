package main

func maximizeSumOfWeights(edges [][]int, k int) int64 {
	n := len(edges) + 1
	dict := make([][][2]int, n)
	out := make([]int, n)
	for _, ed := range edges {
		out[ed[0]] += 1
		out[ed[1]] += 1
		dict[ed[0]] = append(dict[ed[0]], [2]int{ed[1], ed[2]})
		dict[ed[1]] = append(dict[ed[1]], [2]int{ed[0], ed[2]})
	}
	var nodes []int
	for i := 0; i < n; i++ {
		if out[i] <= k {
			continue
		}
		nodes = append(nodes, i)
	}
	return 0
}
