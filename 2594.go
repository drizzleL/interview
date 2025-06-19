package main

import "math"

func repairCars(ranks []int, cars int) int64 {
	check := func(x int) bool {
		var cnt int
		for _, r := range ranks {
			cnt += int(math.Sqrt(float64(x / r)))
		}
		return cnt >= cars
	}
	var l, r int
	for _, rank := range ranks {
		r = max(r, cars*cars*rank)
	}
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return int64(l)
}
