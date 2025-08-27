package main

import "sort"

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	dis := func(a, b int) int {
		return abs(houses[a] - heaters[b])
	}
	check := func(x int) bool {
		for i, j := 0, 0; i < len(houses); i++ {
			for j < len(heaters) && dis(i, j) > x {
				j++
			}
			if j == len(heaters) {
				return false
			}
		}
		return true
	}
	var l, r int
	r = max(dis(len(houses)-1, 0), dis(0, len(heaters)-1))
	for l < r {
		mid := l + (r-l)/2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
