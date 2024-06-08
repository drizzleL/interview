package main

func maximalRectangle(matrix []string) int {
	mm := make([][]int, len(matrix))
	for i := range matrix {
		mm[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		var line int
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				line = 0
				continue
			}
			line++
			mm[i][j] = line
		}
	}
	var ret int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			x := mm[i][j]
			for k := i; k < len(matrix); k++ {
				if mm[k][j] < x {
					x = mm[k][j]
				}
				if x == 0 {
					break
				}
				tmp := (k - i + 1) * x
				if tmp > ret {
					ret = tmp
				}
			}
		}
	}
	return ret
}
