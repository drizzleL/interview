package main

import "math"

func judgeSquareSum(c int) bool {
	m := int(math.Sqrt(float64(c)))
	check := func(b2 int) bool {
		b := int(math.Sqrt(float64(b2)))
		return b*b == b2
	}
	for a := 0; a <= m; a++ {
		if check(c - a*a) {
			return true
		}
	}
	return false
}
