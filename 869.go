package main

import (
	"math"
	"strconv"
)

func reorderedPowerOf2(n int) bool {
	str := strconv.Itoa(n)
	var dict [10]int
	for _, c := range str {
		dict[c-'0'] += 1
	}
	check := func(v int) bool {
		tmp := math.Log2(float64(v))
		return float64(int(tmp)) == tmp
	}
	var helper func(i int, curr int) bool
	helper = func(i int, curr int) bool {
		if i == len(str) {
			return check(curr)
		}
		for j := 0; j < 10; j++ {
			if dict[j] == 0 {
				continue
			}
			if j == 0 && i == 0 {
				continue
			}
			dict[j] -= 1
			if helper(i+1, curr*10+j) {
				return true
			}
			dict[j] += 1
		}
		return false
	}
	return helper(0, 0)
}
