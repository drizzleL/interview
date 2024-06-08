package main

import (
	"math"
)

func isPrintable(targetGrid [][]int) bool {
	posDict := map[int][4]int{}
	for i, line := range targetGrid {
		for j, color := range line {
			if _, ok := posDict[color]; !ok {
				posDict[color] = [4]int{math.MaxInt32, math.MinInt32, math.MaxInt32, math.MinInt32}
			}
			v := [4]int{i, i, j, j}
			v[0] = min(v[0], posDict[color][0])
			v[1] = max(v[1], posDict[color][1])
			v[2] = min(v[2], posDict[color][2])
			v[3] = max(v[3], posDict[color][3])
			posDict[color] = v
		}
	}
	erase := func(color int) {
		for i := posDict[color][0]; i <= posDict[color][1]; i++ {
			for j := posDict[color][2]; j <= posDict[color][3]; j++ {
				targetGrid[i][j] = 0
			}
		}
	}
	check := func(color int) bool {
		for i := posDict[color][0]; i <= posDict[color][1]; i++ {
			for j := posDict[color][2]; j <= posDict[color][3]; j++ {
				if targetGrid[i][j] == 0 || targetGrid[i][j] == color {
					continue
				}
				return false
			}
		}
		return true
	}
	var colors []int
	for color := range posDict {
		colors = append(colors, color)
	}
	for len(colors) != 0 {
		var next []int
		for _, color := range colors {
			if check(color) {
				erase(color)
				continue
			}
			next = append(next, color)
		}
		if len(colors) == len(next) {
			return false
		}
		colors = next
	}
	return true
}
