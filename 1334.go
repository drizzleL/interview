package main

import (
	"math"
)

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	dict := make([][]int, n)
	for i := range dict {
		dict[i] = make([]int, n)
		for j := range dict[i] {
			dict[i][j] = math.MaxInt32
		}
	}
	for i := 0; i < n; i++ {
		dict[i][i] = 0
	}
	for _, ed := range edges {
		dict[ed[0]][ed[1]] = ed[2]
		dict[ed[1]][ed[0]] = ed[2]
	}
	for idx := 0; idx < n; idx++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dict[i][j] = min(dict[i][j], dict[i][idx]+dict[idx][j])
			}
		}
	}
	var ret int
	minVal := math.MaxInt32
	for i := 0; i < n; i++ {
		var cnt int
		for j := 0; j < n; j++ {
			if dict[i][j] <= distanceThreshold {
				cnt += 1
			}
		}
		if cnt <= minVal {
			ret = i
			minVal = cnt
		}
	}
	return ret
}
