package main

import "math"

func shortestPathLength(graph [][]int) int {
	endState := (1 << len(graph)) - 1
	helper := func(i int) int {
		cache := map[[2]int]bool{
			[2]int{i, 0}: true,
		}
		nodes := [][2]int{{i, 0}}
		for step := 0; len(nodes) != 0; step++ {
			var next [][2]int
			for _, n := range nodes {
				i, mask := n[0], n[1]
				for _, child := range graph[i] {
					nextMask := (1 << child) | mask
					if nextMask == endState {
						return step
					}
					if cache[[2]int{child, nextMask}] { // visited
						continue
					}
					cache[[2]int{child, nextMask}] = true
					next = append(next, [2]int{child, nextMask})
				}
			}
			nodes = next
		}
		return 0
	}
	ret := math.MaxInt32
	for i := range graph {
		ret = min(ret, helper(i))
	}
	return ret
}
