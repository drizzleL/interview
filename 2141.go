package main

import "sort"

func maxRunTime(n int, batteries []int) int64 {
	sort.Ints(batteries)
	l, r := 0, 0
	check := func(l, r int, power int) bool {
		return false
	}
	for l < r {
		mid := (l+r)/2 + 1
		if check(0, len(batteries), mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return int64(l)
}
