package main

import (
	"math"
)

func maxDifference3(s string, k int) int {
	ret := math.MinInt32
	getStatus := func(cnta, cntb int) int {
		return ((cnta & 1) << 1) | (cntb & 1)
	}
	base := []byte("01234")
	for _, c1 := range base {
		for _, c2 := range base {
			if c1 == c2 {
				continue
			}
			var status [4]int
			for i := range status {
				status[i] = math.MaxInt32
			}
			var cnta, cntb int
			var precnta, precntb int
			for l, r := 0, 0; r < len(s); r++ {
				if s[r] == c1 {
					cnta += 1
				}
				if s[r] == c2 {
					cntb += 1
				}
				for r-l+1 >= k && cntb-precntb >= 2 {
					prev := getStatus(precnta, precntb)
					status[prev] = min(status[prev], precnta-precntb)
					if s[l] == c1 {
						precnta += 1
					}
					if s[l] == c2 {
						precntb += 1
					}
					l++
				}
				curr := getStatus(cnta, cntb)
				ret = max(ret, cnta-cntb-status[curr^0b10])
			}
		}
	}
	return ret
}
