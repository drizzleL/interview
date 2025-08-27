package main

func areaOfMaxDiagonal(dimensions [][]int) int {
	var ret, diag int
	for i := 0; i < len(dimensions); i++ {
		newDiag := dimensions[i][0]*dimensions[i][0] + dimensions[i][1]*dimensions[i][1]
		if newDiag < diag {
			continue
		}
		area := dimensions[i][0] * dimensions[i][1]
		if newDiag > diag {
			diag = newDiag
			ret = area
			continue
		}
		ret = max(ret, area)
	}
	return ret
}
