package main

import "math"

func countTriples(n int) int {
	var ret int
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			m2 := i*i + j*j
			if m2 > n*n {
				break
			}
			k := int(math.Sqrt(float64(m2)))
			if k*k == m2 {
				if i == j {
					ret += 1
				} else {
					ret += 2
				}
			}
		}
	}
	return ret
}
