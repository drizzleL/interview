package main

import (
	"math"
)

func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	l, r := math.MaxInt32, math.MinInt32
	for _, v := range bloomDay {
		l = min(l, v)
		r = max(r, v)
	}
	for l < r {
		mid := (l + r) / 2
		var flowers, bouquet int
		for _, v := range bloomDay {
			if v > mid {
				flowers = 0
				continue
			}
			flowers += 1
			if flowers >= k {
				bouquet += 1
				flowers -= k
			}
		}
		if bouquet < m {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
