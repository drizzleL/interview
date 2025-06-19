package main

import "sort"

func checkValidCuts(n int, rectangles [][]int) bool {
	checkValid := func(spans [][2]int) bool {
		cnt := 1
		sort.Slice(spans, func(i, j int) bool {
			return spans[i][0] < spans[j][0]
		})
		last := spans[0][1]
		for i := 1; i < len(spans); i++ {
			if spans[i][0] >= last {
				cnt += 1
				last = spans[i][1]
				continue
			}
			last = max(last, spans[i][1])
		}
		return cnt >= 3
	}
	var xline, yLine [][2]int
	for _, rec := range rectangles {
		xline = append(xline, [2]int{rec[0], rec[2]})
		yLine = append(yLine, [2]int{rec[1], rec[3]})
	}
	return checkValid(xline) || checkValid(yLine)
}
