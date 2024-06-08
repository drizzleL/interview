package main

import "math"

func pivotInteger(n int) int {
	tmp := (n*n + n) / 2
	v := int(math.Sqrt(float64(tmp)))
	if v*v == tmp {
		return v
	}
	return -1
}
