package main

import (
	"sort"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	type job struct {
		start  int
		end    int
		profit int
	}
	dict := map[int][]job{} // key is end
	dp := [][2]int{}        // val is [2]int{end, sum}
	dp = append(dp, [2]int{-1, 0})
	for i := 0; i < len(startTime); i++ {
		s, e, p := startTime[i], endTime[i], profit[i]
		if dict[e] == nil {
			dp = append(dp, [2]int{e, 0})
		}
		dict[e] = append(dict[e], job{
			start:  s,
			end:    e,
			profit: p,
		})
	}
	sort.Slice(dp, func(i, j int) bool {
		return dp[i][0] < dp[j][0]
	})
	for i := 1; i < len(dp); i++ {
		dp[i][1] = dp[i-1][1]
		e := dp[i][0]
		for _, job := range dict[e] {
			lastIdx := sort.Search(len(dp), func(i int) bool {
				return dp[i][0] > job.start
			})
			if lastIdx == len(dp) {
				continue
			}
			dp[i][1] = max(dp[i][1], dp[lastIdx-1][1]+job.profit)
		}
	}
	return dp[len(dp)-1][1]
}
