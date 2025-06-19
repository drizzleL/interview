package main

import (
	"sort"
)

func maxTwoEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	beforeVals := [][2]int{{0, 0}}
	for _, ev := range events {
		if ev[2] <= beforeVals[len(beforeVals)-1][1] {
			continue
		}
		if ev[1] == beforeVals[len(beforeVals)-1][0] {
			beforeVals[len(beforeVals)-1][1] = ev[2]
			continue
		}
		beforeVals = append(beforeVals, [2]int{ev[1], ev[2]})
	}
	var ret int
	for _, ev := range events {
		idx := sort.Search(len(beforeVals), func(i int) bool {
			return beforeVals[i][0]+1 > ev[0]
		})
		idx -= 1
		ret = max(ret, ev[2]+beforeVals[idx][1])
	}
	return ret
}
