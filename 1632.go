package main

import (
	"sort"
)

func matrixRankTransform(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	parent := make([]int, m*n)
	for i := range parent {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(a, b int) {
		pa, pb := find(a), find(b)
		if pa > pb {
			pa, pb = pb, pa
		}
		parent[pb] = pa
	}
	rowMax := make([]int, m)
	colMax := make([]int, n)
	var vals [][3]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			vals = append(vals, [3]int{matrix[i][j], i, j})
		}
	}
	sort.Slice(vals, func(i, j int) bool {
		return vals[i][0] < vals[j][0]
	})
	ret := make([][]int, m)
	for i := range ret {
		ret[i] = make([]int, n)
	}
	toGroup := func(nodes [][3]int) [][][3]int {
		row := make([]int, m)
		for i := range row {
			row[i] = -1
		}
		col := make([]int, n)
		for i := range col {
			col[i] = -1
		}
		for _, node := range nodes {
			x, y := node[1], node[2]
			if row[x] != -1 {
				union(x*n+y, x*n+row[x])
			}
			row[x] = y
			if col[y] != -1 {
				union(x*n+y, col[y]*n+y)
			}
			col[y] = x
		}
		groups := map[int][][3]int{}
		for _, node := range nodes {
			x, y := node[1], node[2]
			key := find(x*n + y)
			groups[key] = append(groups[key], node)
		}
		var ret [][][3]int
		for _, v := range groups {
			ret = append(ret, v)
		}
		return ret
	}
	for i := 0; i < len(vals); i++ {
		q := [][3]int{vals[i]}
		v := vals[i][0]
		for i+1 < len(vals) && vals[i+1][0] == v {
			q = append(q, vals[i+1])
			i++
		}
		groups := toGroup(q)
		for _, g := range groups {
			var rank int
			for _, item := range g {
				x, y := item[1], item[2]
				rank = max(rank, rowMax[x]+1)
				rank = max(rank, colMax[y]+1)
			}
			for _, item := range g {
				x, y := item[1], item[2]
				ret[x][y] = rank
				rowMax[x] = rank
				colMax[y] = rank
			}
		}
	}
	return ret
}
