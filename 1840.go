package main

import (
	"sort"
)

func maxBuilding(n int, restrictions [][]int) int {
	restrictions = append(restrictions, []int{1, 0}, []int{n, n - 1})
	sort.Slice(restrictions, func(i, j int) bool {
		return restrictions[i][0] < restrictions[j][0]
	})
	for i := 1; i < len(restrictions); i++ {
		restrictions[i][1] = min(restrictions[i][1], restrictions[i-1][1]+restrictions[i][0]-restrictions[i-1][0])
	}
	for i := len(restrictions) - 2; i >= 0; i-- {
		restrictions[i][1] = min(restrictions[i][1], restrictions[i+1][1]+restrictions[i+1][0]-restrictions[i][0])
	}
	var ret int
	for i := 1; i < len(restrictions); i++ {
		a, b := restrictions[i-1], restrictions[i]
		ret = max(ret, max(a[1], b[1])+(b[0]-a[0]+abs(b[1]-a[1]))/2)
	}
	return ret

}
