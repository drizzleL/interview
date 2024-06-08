package main

import "sort"

func kClosest(points [][]int, k int) [][]int {
	if k >= len(points) {
		return points
	}
	type pointWithDistance struct {
		point    []int
		distance int
	}
	pp := make([]pointWithDistance, len(points))
	for i := range points {
		pp[i] = pointWithDistance{
			point:    points[i],
			distance: points[i][0]*points[i][0] + points[i][1]*points[i][1],
		}
	}
	sort.Slice(pp, func(i, j int) bool {
		return pp[i].distance < pp[j].distance
	})
	ret := make([][]int, 0, k)
	for i := 0; i < k; i++ {
		ret = append(ret, pp[i].point)
	}
	return ret
}
