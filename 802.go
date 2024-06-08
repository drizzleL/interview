package main

import (
	"sort"
)

func eventualSafeNodes(graph [][]int) []int {
	seen := make([]bool, len(graph))
	backDict := map[int][]int{}
	nodes := []int{}
	out := make([]int, len(graph))
	var ret []int
	for i, g := range graph {
		out[i] = len(g)
		if len(g) == 0 {
			nodes = append(nodes, i)
			ret = append(ret, i)
			seen[i] = true
			continue
		}
		for _, next := range g {
			backDict[next] = append(backDict[next], i)
		}
	}
	for len(nodes) != 0 {
		var next []int
		for _, n := range nodes {
			for _, b := range backDict[n] {
				out[b] -= 1
				if out[b] == 0 {
					ret = append(ret, b)
					next = append(next, b)
				}
			}
		}
		nodes = next
	}
	sort.Ints(ret)
	return ret
}
