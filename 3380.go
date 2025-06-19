package main

import (
	"sort"
)

func maxRectangleArea(points [][]int) int {
	xDict, yDict := map[int][]int{}, map[int][]int{}
	dict := map[[2]int]int{}
	for i, point := range points {
		xDict[point[0]] = append(xDict[point[0]], i)
		yDict[point[1]] = append(yDict[point[1]], i)
		dict[[2]int{point[0], point[1]}] = i
	}
	for _, v := range xDict {
		sort.Slice(v, func(i, j int) bool {
			idx1, idx2 := v[i], v[j]
			return points[idx1][1] < points[idx2][1]
		})
	}
	for _, v := range yDict {
		sort.Slice(v, func(i, j int) bool {
			idx1, idx2 := v[i], v[j]
			return points[idx1][0] < points[idx2][0]
		})
	}
	helper := func(i int) int {
		sameX := xDict[points[i][0]]
		nextX := sort.Search(len(sameX), func(j int) bool {
			return points[sameX[j]][1] >= points[i][1]
		}) + 1
		if nextX == len(sameX) {
			return -1
		}
		sameY := yDict[points[i][1]]
		nextY := sort.Search(len(sameY), func(j int) bool {
			return points[sameY[j]][0] >= points[i][0]
		}) + 1
		if nextY == len(sameY) {
			return -1
		}
		a, b := points[sameX[nextX]][1]-points[i][1], points[sameY[nextY]][0]-points[i][0]
		idx3, ok := dict[[2]int{points[i][0] + b, points[i][1] + a}]
		if !ok {
			return -1
		}
		found := map[int]bool{
			i:            true,
			sameX[nextX]: true,
			sameY[nextY]: true,
			idx3:         true,
		}
		for j, point := range points {
			if found[j] {
				continue
			}
			if point[0] < points[i][0] || point[0] > points[i][0]+b {
				continue
			}
			if point[1] < points[i][1] || point[1] > points[i][1]+a {
				continue
			}
			return -1
		}
		return a * b
	}
	ret := -1
	for i := range points {
		if x := helper(i); x != -1 {
			ret = max(ret, x)
		}
	}
	return ret
}
