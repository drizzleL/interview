package main

import "math"

func minAreaRect(points [][]int) int {
	ret := math.MaxInt32
	dict := map[[2]int]bool{}
	for _, p := range points {
		dict[[2]int{p[0], p[1]}] = true
	}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]
			if p1[0] == p2[0] || p1[1] == p2[1] {
				continue
			}
			p3, p4 := [2]int{p1[0], p2[1]}, [2]int{p2[0], p1[1]}
			if !dict[p3] || !dict[p4] {
				continue
			}
			ret = min(ret, abs(p1[0]-p2[0])*abs(p1[1]-p2[1]))
		}
	}
	if ret == math.MaxInt32 {
		return 0
	}
	return ret
}
