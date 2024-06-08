package main

import "math"

func minSpeedOnTime(dist []int, hour float64) int {
	if math.Ceil(hour) < float64(len(dist)) {
		return -1
	}
	canReach := func(speed int) bool {
		var h float64
		for i, d := range dist {
			takes := float64(d) / float64(speed)
			if i == len(dist)-1 { // last, can float
				h += takes
			} else {
				h += math.Ceil(takes)
			}
		}
		return h <= hour
	}
	l, r := 0, int(1e7)
	for l < r {
		mid := (l + r) / 2
		if !canReach(mid) {
			l = mid + 1
			continue
		}
		r = mid
	}
	return l
}
