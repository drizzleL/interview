package main

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	if start_node == end_node {
		return 1
	}
	dict := map[int][][2]int{}
	for i, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], [2]int{ed[1], i})
		dict[ed[1]] = append(dict[ed[1]], [2]int{ed[0], i})
	}
	nodes := []int{start_node}
	dp := make([]float64, n)
	dp[start_node] = 1
	for len(nodes) != 0 {
		var newnodes []int
		for _, node := range nodes {
			for _, child := range dict[node] {
				tmp := succProb[child[1]] * dp[node]
				if tmp > dp[child[0]] {
					if child[0] != end_node {
						newnodes = append(newnodes, child[0])
					}
					dp[child[0]] = tmp
				}
			}
		}
		nodes = newnodes
	}
	return dp[end_node]
}

func maxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
