package main

import (
	"math"
)

func earliestFinishTime2(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	var landline [][]int
	for i := range landStartTime {
		landline = append(landline, []int{landStartTime[i], landDuration[i]})
	}
	var waterline [][]int
	for i := range waterStartTime {
		waterline = append(waterline, []int{waterStartTime[i], waterDuration[i]})
	}
	check := func(line1, line2 [][]int) int {
		minEnd := math.MaxInt32
		for i := range line1 {
			minEnd = min(minEnd, line1[i][0]+line1[i][1])
		}
		ret := math.MaxInt32
		for i := range line2 {
			start, takes := line2[i][0], line2[i][1]
			if start >= minEnd {
				ret = min(ret, start+takes)
			} else {
				ret = min(ret, minEnd+takes)
			}
		}
		return ret
	}
	return min(check(landline, waterline), check(waterline, landline))
}
