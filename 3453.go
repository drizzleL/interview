package main

import (
	"math"
)

func separateSquares(squares [][]int) float64 {
	var sum float64
	for _, sq := range squares {
		sum += float64(sq[2] * sq[2])
	}
	sum /= 2
	getBelowSum := func(y float64) float64 {
		var belowSum float64
		for _, sq := range squares {
			if y <= float64(sq[1]) {
				continue
			}
			y2 := y
			if y2 > float64(sq[1]+sq[2]) {
				y2 = float64(sq[1] + sq[2])
			}
			belowSum += (y2 - float64(sq[1])) * float64(sq[2])
		}
		return belowSum
	}
	lInt, rInt := math.MaxInt64, math.MinInt64
	for _, sq := range squares {
		lInt = min(lInt, sq[1])
		rInt = max(rInt, sq[1]+sq[2])
	}
	l, r := float64(lInt), float64(rInt)
	for i := 0; i < 60; i++ {
		mid := (l + r) / 2
		belowSum := getBelowSum(mid)
		if belowSum >= sum {
			r = mid
		} else {
			l = mid
		}
	}
	return l
}
