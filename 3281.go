package main

import "sort"

func maxPossibleScore(start []int, d int) int {
	sort.Ints(start)
	var l, r int
	for i := 1; i < len(start); i++ {
		gap := start[i] - start[i-1]
		l = min(l, max(0, gap-d))
		r = max(r, gap+d)
	}
	check := func(x int) bool {
		last := start[0]
		for i := 1; i < len(start); i++ {
			if start[i]+d-last < x {
				return false
			}
			last = max(last+x, start[i])
		}
		return true
	}
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}
