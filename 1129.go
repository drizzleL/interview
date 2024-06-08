package main

import "math"

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	blueDict := make([][]int, n)
	redDict := make([][]int, n)
	for _, edge := range redEdges {
		redDict[edge[0]] = append(redDict[edge[0]], edge[1])
	}
	for _, edge := range blueEdges {
		blueDict[edge[0]] = append(blueDict[edge[0]], edge[1])
	}
	dicts := make([][][]int, 2)
	dicts[0] = redDict
	dicts[1] = blueDict

	ret := make([]int, n)
	for i := range ret {
		ret[i] = math.MaxInt32
	}

	helper := func(dictIdx int, nodes []int) {
		visiteds := make([][]bool, 2)
		visiteds[0] = make([]bool, n)
		visiteds[1] = make([]bool, n)
		var step, visitedIdx int
		for len(nodes) != 0 {
			var nextnodes []int
			dict := dicts[dictIdx]
			visited := visiteds[visitedIdx]
			for _, node := range nodes {
				ret[node] = min(ret[node], step)
				for _, next := range dict[node] {
					if visited[next] {
						continue
					}
					visited[next] = true
					nextnodes = append(nextnodes, next)
				}
			}
			dictIdx = 1 - dictIdx
			visitedIdx = 1 - visitedIdx
			step++
			nodes = nextnodes
		}

	}

	helper(0, []int{0})
	helper(1, []int{0})

	for i, v := range ret {
		if v == math.MaxInt32 {
			ret[i] = -1
		}
	}
	return ret
}
