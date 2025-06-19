package main

import (
	"sort"
)

func eventualSafeNodes2(graph [][]int) []int {
	safe := make([]int, len(graph))
	for i, node := range graph {
		if len(node) == 0 {
			safe[i] = 1
		}
	}
	var helper func(i int) (ret int)
	helper = func(i int) (ret int) {
		if safe[i] != 0 {
			return safe[i]
		}
		defer func() {
			safe[i] = ret
		}()
		safe[i] = 2
		ret = 1
		for _, next := range graph[i] {
			if helper(next) == 2 {
				ret = 2
				break
			}
		}
		return
	}
	for i := range graph {
		safe[i] = helper(i)
	}
	var ret []int
	for i, v := range safe {
		if v == 1 {
			ret = append(ret, i)
		}
	}
	return ret
}
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
