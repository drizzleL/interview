package main

func maxSum6(grid [][]int) int {
	var ret int
	sums := make([][]int, len(grid))
	for i, row := range grid {
		sums[i] = make([]int, len(row))
		for j, num := range row {
			if j != 0 {
				sums[i][j] = sums[i][j-1]
			}
			sums[i][j] += num
		}
	}
	query := func(row int, l, r int) int {
		ret := sums[row][r]
		if l != 0 {
			ret -= sums[row][l-1]
		}
		return ret
	}
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			tmp := query(i-1, j-1, j+1) + query(i+1, j-1, j+1) + query(i, j-1, j+1) - grid[i][j-1] - grid[i][j+1]
			ret = max(ret, tmp)
		}
	}
	return ret
}
