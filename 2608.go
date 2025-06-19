package main

import "math"

func findShortestCycle(n int, edges [][]int) int {
	dict := make([][]int, n)
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], ed[1])
		dict[ed[1]] = append(dict[ed[1]], ed[0])
	}
	helper := func(i int) int {
		nodes := [][2]int{{i, -1}}
		seen := make([]bool, n)
		ret := math.MaxInt32
		for step := 0; ret != math.MaxInt32 && len(nodes) > 0; step++ {
			var next [][2]int
			nextSeen := make([]bool, n)
			for _, node := range nodes {
				if seen[node[0]] {
					return step * 2
				}
				seen[node[0]] = true
				for _, child := range dict[node[0]] {
					if node[1] == child {
						continue
					}
					if nextSeen[child] {
						ret = min(ret, (step+1)*2)
						continue
					}
					nextSeen[child] = true
					if seen[child] {
						ret = min(ret, step*2+1)
						continue
					}
					next = append(next, [2]int{child, node[0]})
				}
			}
			nodes = next
		}
		return ret
	}
	ret := -1
	for i := 0; i < n; i++ {
		tmp := helper(i)
		if tmp == math.MaxInt32 {
			continue
		}
		if ret == -1 || tmp < ret {
			ret = tmp
		}
	}
	return ret
}
