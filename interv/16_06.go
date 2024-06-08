package main

import (
	"math"
	"sort"
)

func smallestDifference(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	ret := math.MaxInt32
	for i, j := 0, 0; i < len(a) && j < len(b); {
		v1, v2 := a[i], b[j]
		if v1 == v2 {
			return 0
		}
		if v1 > v2 {
			ret = min(ret, v1-v2)
			i++
		} else {
			ret = min(ret, v2-v1)
			j++
		}
	}
	return ret
}
