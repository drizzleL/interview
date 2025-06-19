package main

import (
	"math"
)

func findNumbers(nums []int) int {
	var ret int
	for _, num := range nums {
		ret += int(math.Log10(float64(num))) % 2
	}
	return ret
}
