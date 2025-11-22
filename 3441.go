package main

import (
	"fmt"
	"math"
)

func minCostGoodCaption(caption string) string {
	if len(caption) <= 2 {
		return ""
	}
	dp := make([][][]int, len(caption)+1)
	for i := range dp {
		dp[i] = make([][]int, 26)
		for j := range dp[i] {
			dp[i][j] = make([]int, 3)
		}
	}
	for j := 0; j < 26; j++ {
		for k := 0; k < 3; k++ {
			dp[0][j][k] = math.MaxInt32
		}
	}
	var best int
	for i := 0; i < len(caption); i++ {
		c := caption[len(caption)-i-1]
		newbest := math.MaxInt32
		for j := 0; j < 26; j++ {
			curr := abs(int(c-'a') - j)
			dp[i+1][j][0] = best + curr
			dp[i+1][j][1] = dp[i][j][0] + curr
			dp[i+1][j][2] = min(dp[i][j][1], dp[i][j][2]) + curr
			newbest = min(newbest, dp[i+1][j][2])
		}
		best = newbest
	}
	find := func(i int, best int) int {
		var ret int
		for j := 25; j >= 0; j-- {
			if dp[i][j][2] == best {
				ret = j
			}
		}
		return ret
	}
	var b []byte
	for i := len(caption); i >= 1; i-- {
		j := find(i, best)
		c := caption[len(caption)-i]
		diff := abs(int(c-'a') - j)
		fmt.Println(best, string(c), diff)
		best -= diff
		b = append(b, byte('a'+j))
	}
	return string(b)
}
