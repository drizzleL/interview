package main

import (
	"math"
)

func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
	dp := make([]int, numLaps+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for _, tire := range tires {
		a, b := tire[0], tire[1]
		var last int
		for i := 1; i <= numLaps; i++ {
			dp[i] = min(dp[i], last+a)
			last += a
			a *= b
			if a > math.MaxInt32 {
				break
			}
		}
	}
	for i := 2; i <= numLaps; i++ {
		for j := 1; j*2 <= i; j++ {
			dp[i] = min(dp[i], dp[j]+dp[i-j]+changeTime)
		}
	}
	return dp[numLaps]
}
