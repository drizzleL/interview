package main

import "math"

func minimumTime(time []int, totalTrips int) int64 {
	minT, maxT := math.MaxInt32, math.MinInt32
	for _, t := range time {
		currT := t * totalTrips
		minT = min(minT, maxT)
		maxT = max(maxT, currT)
	}
	for minT < maxT {
		m := (minT + maxT) / 2
		var trips int
		for _, t := range time {
			trips += m / t
		}
		if trips < totalTrips {
			minT = m + 1
		} else {
			maxT = m
		}
	}
	return int64(minT)
}
