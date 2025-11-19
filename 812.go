package main

import "math"

func largestTriangleArea(points [][]int) float64 {
	var ret float64
	for i := 0; i < len(points); i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < len(points); j++ {
			x2, y2 := points[j][0], points[j][1]
			for k := j + 1; k < len(points); k++ {
				x3, y3 := points[k][0], points[k][1]
				area := math.Abs(0.5 * float64(x1*(y2-y3)+x2*(y3-y1)+x3*(y1-y2)))
				ret = math.Max(ret, area)
			}
		}
	}
	return ret
}
