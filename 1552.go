package main

import (
	"sort"
)

func maxDistance(position []int, m int) int {
	sort.Ints(position)
	l := 1
	r := (position[len(position)-1] - position[0]) / (m - 1)
	check := func(gap int) bool {
		left := m - 1
		idx := 0
		for left > 0 {
			idx = sort.SearchInts(position, position[idx]+gap)
			if idx >= len(position) {
				return false
			}
			left -= 1
		}
		return true
	}
	for l < r {
		mid := (l+r)/2 + 1
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}
