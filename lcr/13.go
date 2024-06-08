package main

type NumMatrix struct {
	data [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sum := make([][]int, len(matrix))
	for i := range sum {
		sum[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		var lineaccum int
		for j := 0; j < len(matrix[0]); j++ {
			lineaccum += matrix[i][j]
			sum[i][j] = lineaccum
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
		}
	}
	return NumMatrix{
		data: sum,
	}
}

func (this *NumMatrix) sum(row int, col int) int {
	if row < 0 || col < 0 {
		return 0
	}
	return this.data[row][col]
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sum(row2, col2) - this.sum(row1-1, col2) - this.sum(row2, col1-1) + this.sum(row1-1, col1-1)
}
