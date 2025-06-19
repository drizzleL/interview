package main

import (
	"log"
	"math"
)

func minimizedMaximum(n int, quantities []int) int {
	var r int
	for _, q := range quantities {
		r = max(r, q)
	}
	l := 1
	check := func(x int) bool {
		var cnt int
		for _, q := range quantities {
			cnt += int(math.Ceil(float64(q) / float64(x)))
		}
		log.Println(x, cnt)
		return cnt <= n
	}
	for l < r {
		mid := (l + r) / 2
		log.Println(l, r, mid, check(mid))
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
