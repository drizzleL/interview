package main

func findFarmland(land [][]int) [][]int {
	var ret [][]int
	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[0]); j++ {
			if land[i][j] != 1 {
				continue
			}
			r, d := j, i
			for ; r < len(land[0]) && land[i][r] == 1; r++ {
			}
			for ; d < len(land) && land[d][r-1] == 1; d++ {
			}
			for i2 := i; i2 < d; i2++ {
				for j2 := j; j2 < r; j2++ {
					land[i2][j2] = 2
				}
			}
			ret = append(ret, []int{i, j, d - 1, r - 1})
		}
	}
	return ret
}
