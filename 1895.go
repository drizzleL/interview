package main

func largestMagicSquare(grid [][]int) int {
	rowSums := make([][]int, len(grid)+1)
	for i := range rowSums {
		rowSums[i] = make([]int, len(grid[0])+1)
	}
	colSums := make([][]int, len(grid)+1)
	for i := range colSums {
		colSums[i] = make([]int, len(grid[0])+1)
	}
	for i := range grid {
		for j := range grid[i] {
			rowSums[i+1][j+1] += rowSums[i][j+1] + grid[i][j]
			colSums[i+1][j+1] += colSums[i+1][j] + grid[i][j]
		}
	}
	check := func(i, j int, k int) bool {
		baseSum := rowSums[i+k][j+1] - rowSums[i][j+1]
		for j2 := j + 1; j2 <= j+k-1; j2++ {
			if rowSums[i+k][j2+1]-rowSums[i][j2+1] != baseSum {
				return false
			}
		}
		for i2 := i; i2 <= i+k-1; i2++ {
			if colSums[i2+1][j+k]-colSums[i2+1][j] != baseSum {
				return false
			}
		}
		var s1, s2 int
		for m := 0; m < k; m++ {
			s1 += grid[i+m][j+m]
			s2 += grid[i+m][j+k-1-m]
		}
		return s1 == baseSum && s2 == baseSum
	}
	ret := 1
	for i := range grid {
		for j := range grid[i] {
			maxLimit := min(len(grid)-i, len(grid[0])-j)
			for k := ret + 1; k <= maxLimit; k++ {
				if check(i, j, k) {
					ret = max(ret, k)
				}
			}
		}
	}
	return ret
}
