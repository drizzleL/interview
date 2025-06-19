package main

import (
	"math"
)

func findMaximumNumber(k int64, x int) int64 {
	dp := []int{0}
	for dp[len(dp)-1] < int(k) {
		v := dp[len(dp)-1]
		v *= 2
		if (len(dp))%x == 0 {
			v += 1 << (len(dp) - 1)
		}
		dp = append(dp, v)
	}
	var helper func(v int) int
	helper = func(v int) int {
		if v == 0 {
			return 0
		}
		bit := int(math.Log2(float64(v)))
		ret := dp[bit]
		if (bit+1)%x == 0 {
			ret += v - 1<<bit + 1
		}
		ret += helper(v - 1<<bit)
		return ret
	}
	l, r := 0, 1<<len(dp)
	for l < r {
		mid := (l + r + 1) / 2
		if helper(mid) > int(k) {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return int64(l)
}
