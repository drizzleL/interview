package main

func countVisitedNodes(edges []int) []int {
	n := len(edges)
	seen := make([]int, n)
	for i := range seen {
		seen[i] = -1
	}
	var cycleSizes []int
	var cycleNodes []int
	for i := 0; i < n; i++ {
		if seen[i] != -1 {
			continue
		}
		node := i
		for seen[node] == -1 {
			seen[node] = i
			node = edges[node]
		}
		if seen[node] != i {
			continue
		}
		cycleNodes = append(cycleNodes, node)
	}
	cycleDict := make([]int, n)
	for i := range cycleDict {
		cycleDict[i] = -1
	}
	seen2 := make([]bool, n) // reset seen
	for _, cycleNode := range cycleNodes {
		var cycleSize int
		for !seen2[cycleNode] {
			seen2[cycleNode] = true
			cycleDict[cycleNode] = len(cycleSizes)
			cycleNode = edges[cycleNode]
			cycleSize += 1
		}
		cycleSizes = append(cycleSizes, cycleSize)
	}
	ret := make([]int, n)
	var dfs func(i int) int
	dfs = func(i int) int {
		if ret[i] != 0 {
			return ret[i]
		}
		cycleId := cycleDict[i]
		if cycleId != -1 {
			return cycleSizes[cycleId]
		}
		ret[i] = 1 + dfs(edges[i])
		return ret[i]
	}
	for i := 0; i < n; i++ {
		ret[i] = dfs(i)
	}
	return ret
}
