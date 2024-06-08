package main

import "sort"

func pileBox(box [][]int) int {
	sort.Slice(box, func(i, j int) bool {
		return box[i][0] < box[j][0]
	})
	dp := make([]int, len(box))
	for i := 0; i < len(box); i++ {
		a := box[i]
		dp[i] = a[2]
		var bigger []int
		for j := 0; j < i; j++ {
			b := box[j]
			switch {
			case a[0] > b[0] && a[1] > b[1] && a[2] > b[2]:
				dp[i] = max(dp[i], dp[j]+a[2])
			case a[0] < b[0] && a[1] < b[1] && a[2] < b[2]:
				bigger = append(bigger, j)
			}
		}
		for _, biggerIdx := range bigger {
			dp[biggerIdx] = max(dp[biggerIdx], dp[i]+box[biggerIdx][2])
		}
	}
	var ret int
	for _, v := range dp {
		ret = max(ret, v)
	}
	return ret
}
