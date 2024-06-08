package main

import "math"

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	m := len(triangle)
	for i := 1; i < m; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
			} else if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				triangle[i][j] += min(triangle[i-1][j], triangle[i-1][j-1])
			}
		}
	}
	ret := math.MaxInt32
	for j := 0; j < len(triangle[m-1]); j++ {
		if triangle[m-1][j] < ret {
			ret = triangle[m-1][j]
		}
	}
	return ret
}
