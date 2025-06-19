package main

import "math"

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	dict := make([][]int, len(edges)+1)
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], ed[1])
		dict[ed[1]] = append(dict[ed[1]], ed[0])
	}
	amount[bob] = 0
	bobSeen := make([]bool, len(edges)+1)
	bobSeen[bob] = true
	var getBobPath func(start int, path []int) []int
	getBobPath = func(start int, path []int) []int {
		if start == 0 {
			return path
		}
		for _, child := range dict[start] {
			if bobSeen[child] {
				continue
			}
			bobSeen[child] = true
			path2 := getBobPath(child, append(path, child))
			if len(path2) != 0 {
				return path2
			}
		}
		return nil
	}
	bobPath := getBobPath(bob, nil)
	amount[bob] = 0
	nodes := [][2]int{{0, amount[0]}}
	seen := make([]bool, len(edges)+1)
	seen[0] = true
	ret := math.MinInt32
	for step := 0; len(nodes) != 0; step++ {
		var next [][2]int
		for _, node := range nodes {
			var flag bool
			for _, child := range dict[node[0]] {
				if seen[child] {
					continue
				}
				flag = true
				seen[child] = true
				extra := amount[child]
				if step < len(bobPath) && bobPath[step] == child {
					extra /= 2
				}
				next = append(next, [2]int{child, node[1] + extra})
			}
			if !flag { // no child found
				ret = max(ret, node[1])
			}
		}
		if step < len(bobPath) {
			amount[bobPath[step]] = 0 // reset bob path's amount
		}
		nodes = next
	}
	return ret
}
