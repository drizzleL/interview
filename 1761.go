package main

import "math"

func minTrioDegree(n int, edges [][]int) int {
	degree := make([]int, n)
	dict := map[[2]int]bool{}
	for _, ed := range edges {
		degree[ed[0]-1] += 1
		degree[ed[1]-1] += 1
		dict[[2]int{ed[0] - 1, ed[1] - 1}] = true
		dict[[2]int{ed[1] - 1, ed[0] - 1}] = true
	}
	ret := math.MaxInt32
	for _, ed := range edges {
		a, b := ed[0]-1, ed[1]-1
		for c := 0; c < n; c++ {
			if a == c || b == c {
				continue
			}
			if dict[[2]int{a, c}] && dict[[2]int{b, c}] {
				ret = min(ret, degree[a]+degree[b]+degree[c]-6)
			}
		}
	}
	if ret == math.MaxInt32 {
		return -1
	}
	return ret
}
