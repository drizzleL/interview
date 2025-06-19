package main

func maxEqualRowsAfterFlips(matrix [][]int) int {
	dict := map[string]int{}
	for i := 0; i < len(matrix); i++ {
		flag := matrix[i][0]
		var c byte
		var b []byte
		for j := 0; j < len(matrix[0]); j++ {
			if flag == 1 {
				matrix[i][j] = 1 - matrix[i][j]
			}
			if j != 0 && j%8 == 0 {
				b = append(b, c)
				c = 0
			}
			c <<= 1
			c |= byte(matrix[i][j])
		}
		b = append(b, c)
		dict[string(b)] += 1
	}
	var ret int
	for _, v := range dict {
		ret = max(ret, v)
	}
	return ret
}
