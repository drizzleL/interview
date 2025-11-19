package main

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	for i := 1; i < n; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
				continue
			}
			if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][j-1]
				continue
			}
			triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
		}
	}
	ret := triangle[n-1][0]
	for j := 0; j < len(triangle[n-1]); j++ {
		ret = min(ret, triangle[n-1][j])
	}
	return ret
}
