package main

import (
	"math"
	"sort"
)

func minCost19(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	var cells [][3]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cells = append(cells, [3]int{i, j, grid[i][j]})
		}
	}
	sort.Slice(cells, func(i, j int) bool {
		return cells[i][2] < cells[j][2]
	})
	valDict := map[int]int{}
	for i := len(cells) - 1; i >= 0; i-- {
		if _, ok := valDict[cells[i][2]]; ok {
			continue
		}
		valDict[cells[i][2]] = i
	}
	costs := make([][][]int, k+1)
	for i := range costs {
		costs[i] = make([][]int, m)
		for j := range costs[i] {
			costs[i][j] = make([]int, n)
			for m := range costs[i][j] {
				costs[i][j][m] = math.MaxInt32
			}
		}
		costs[i][m-1][n-1] = 0
	}
	var valCost [][2]int
	for l := 0; l <= k; l++ {
		for i := m - 1; i >= 0; i-- {
			for j := n - 1; j >= 0; j-- {
				if i != m-1 {
					costs[l][i][j] = min(costs[l][i][j], grid[i+1][j]+costs[l][i+1][j])
				}
				if j != n-1 {
					costs[l][i][j] = min(costs[l][i][j], grid[i][j+1]+costs[l][i][j+1])
				}
				if l != 0 {
					idx := valDict[grid[i][j]]
					costs[l][i][j] = min(costs[l][i][j], valCost[idx][1])
				}
			}
		}
		var newValCost [][2]int
		lastCost := math.MaxInt32
		for _, cell := range cells {
			i, j := cell[0], cell[1]
			lastCost = min(lastCost, costs[l][i][j])
			newValCost = append(newValCost, [2]int{cell[2], lastCost})
		}
		valCost = newValCost
	}
	ret := math.MaxInt32
	for l := 0; l <= k; l++ {
		ret = min(ret, costs[l][0][0])
	}
	return ret
}
