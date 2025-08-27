package main

import "math"

func maximumEnergy(energy []int, k int) int {
	ret := math.MinInt32
	for i := len(energy) - 1; i >= len(energy)-k; i-- {
		var v int
		for j := i; j >= 0; j -= k {
			v += energy[j]
			ret = max(ret, v)
		}
	}
	return ret
}
