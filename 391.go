package main

import (
	"math"
)

func isRectangleCover(rectangles [][]int) bool {
	dict := map[[2]int]int{}
	l, r, up, bot := math.MaxInt32, math.MinInt32, math.MinInt32, math.MaxInt32
	var sum int
	for _, rec := range rectangles {
		x1, y1, x2, y2 := rec[0], rec[1], rec[2], rec[3]
		l = min(l, x1)
		r = max(r, x2)
		bot = min(bot, y1)
		up = max(up, y2)
		dict[[2]int{x1, y1}] += 1
		dict[[2]int{x1, y2}] += 1
		dict[[2]int{x2, y1}] += 1
		dict[[2]int{x2, y2}] += 1
		sum += (x2 - x1) * (y2 - y1)
	}
	if sum != (r-l)*(up-bot) {
		return false
	}
	for _, x := range []int{l, r} {
		for _, y := range []int{up, bot} {
			if dict[[2]int{x, y}] != 1 {
				return false
			}
		}
	}
	var cnt int
	for _, v := range dict {
		if v == 1 {
			cnt += 1
			continue
		}
		if v%2 == 1 {
			return false
		}
	}
	return cnt == 4
}
