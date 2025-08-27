package main

import "math"

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	var landline [][]int
	for i := range landStartTime {
		landline = append(landline, []int{landStartTime[i], landDuration[i]})
	}
	var waterline [][]int
	for i := range waterStartTime {
		waterline = append(waterline, []int{waterStartTime[i], waterDuration[i]})
	}
	check := func(line1, line2 [][]int) int {
		ret := math.MaxInt32
		for i := range line1 {
			start, takes := line1[i][0], line1[i][1]
			end := start + takes
			for j := range line2 {
				start2, takes2 := line2[j][0], line2[j][1]
				if start2 >= end {
					ret = min(ret, start2+takes2)
				} else {
					ret = min(ret, end+takes2)
				}
			}
		}
		return ret
	}
	return min(check(landline, waterline), check(waterline, landline))
}
