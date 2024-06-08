package main

import (
	"strconv"
)

func findKthNumber(k int) int {
	size := 9
	start := 1
	bits := 1
	for {
		if k <= size*bits {
			break
		}
		k -= size * bits
		start *= 10
		size *= 10
		bits++
	}
	a, b := (k-1)/bits, (k-1)%bits
	v := start + a
	return int(strconv.Itoa(v)[b] - '0')
}
