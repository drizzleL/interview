package main

import "sort"

func maxArea(coords [][]int) int64 {
	xDict := map[int][2]int{}
	yDict := map[int][2]int{}
	for _, cor := range coords {
		if old, ok := xDict[cor[0]]; ok {
			xDict[cor[0]] = [2]int{min(old[0], cor[1]), max(old[1], cor[1])}
		} else {
			xDict[cor[0]] = [2]int{cor[1], cor[1]}
		}
		if old, ok := yDict[cor[1]]; ok {
			yDict[cor[1]] = [2]int{min(old[0], cor[0]), max(old[1], cor[0])}
		} else {
			yDict[cor[1]] = [2]int{cor[0], cor[0]}
		}
	}
	ret := -1
	helper := func(dict map[int][2]int) {
		var points [][3]int
		for k, v := range dict {
			points = append(points, [3]int{k, v[0], v[1]})
		}
		sort.Slice(points, func(i, j int) bool {
			return points[i][0] < points[j][0]
		})
		for i := 0; i < len(points); i++ {
			p := points[i]
			if p[1] == p[2] {
				continue
			}
			l := max(p[0]-points[0][0], points[len(points)-1][0]-p[0])
			if l == 0 {
				continue
			}
			ret = max(ret, (p[2]-p[1])*l)
		}
	}
	helper(xDict)
	helper(yDict)
	return int64(ret)
}
