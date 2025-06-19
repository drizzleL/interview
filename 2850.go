package main

import "math"

func minimumMoves(grid [][]int) int {
	var arr1 [][2]int
	var arr2 [][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i][j] == 1 {
				continue
			}
			if grid[i][j] == 0 {
				arr1 = append(arr1, [2]int{i, j})
				continue
			}
			arr2 = append(arr2, [3]int{i, j, grid[i][j] - 1})
		}
	}
	dist := func(i, j, i2, j2 int) int {
		return abs(i-i2) + abs(j-j2)
	}
	var helper func(i int) int
	helper = func(i int) (ret int) {
		if i == len(arr1) {
			return 0
		}
		ret = math.MaxInt32
		for j, v := range arr2 {
			if v[2] == 0 {
				continue
			}
			arr2[j][2] -= 1
			ret = min(ret, helper(i+1)+dist(arr1[i][0], arr1[i][1], v[0], v[1]))
			arr2[j][2] += 1
		}
		return
	}
	return helper(0)
}
