package main

func luckyNumbers(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	maxRowIdx := make([]int, len(matrix))
	minColIdx := make([]int, len(matrix[0]))
	for i, line := range matrix {
		for j, num := range line {
			if num < line[maxRowIdx[i]] {
				maxRowIdx[i] = j
			}
			if num > matrix[minColIdx[j]][j] {
				minColIdx[j] = i
			}
		}
	}
	var ret []int
	for i, j := range maxRowIdx {
		if minColIdx[j] == i {
			ret = append(ret, matrix[i][j])
		}
	}
	return ret
}
