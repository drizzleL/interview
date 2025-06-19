package main

import (
	"math"
	"math/bits"
)

func uniqueXorTriplets(nums []int) int {
	size := len(nums)
	switch size {
	case 1, 2:
		return size
	default:
		bitLen := bits.Len(uint(size))
		return int(math.Pow(2, float64(bitLen)+1))
	}
}
