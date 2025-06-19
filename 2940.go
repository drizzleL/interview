package main

import (
	"sort"
)

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	ret := make([]int, len(queries))
	var qq [][]int
	for i, q := range queries {
		a, b := q[0], q[1]
		if a == b {
			ret[i] = a
			continue
		}
		if a > b {
			a, b = b, a
		}
		if heights[b] > heights[a] {
			ret[i] = b
			continue
		}
		qq = append(qq, []int{i, b, heights[a] + 1})
	}
	sort.Slice(qq, func(i, j int) bool {
		return qq[i][1] < qq[j][1]
	})
	var h [][2]int
	for i, j := len(heights)-1, len(qq)-1; i >= 0 && j >= 0; i-- {
		for j >= 0 && qq[j][1] == i {
			q := qq[j]
			idx := sort.Search(len(h), func(k int) bool {
				return h[k][1] < q[2]
			}) - 1
			if idx == -1 {
				ret[q[0]] = -1
			} else {
				ret[q[0]] = h[idx][0]
			}
			j--
		}
		for len(h) != 0 && h[len(h)-1][1] <= heights[i] {
			h = h[:len(h)-1]
		}
		h = append(h, [2]int{i, heights[i]})
	}
	return ret
}
