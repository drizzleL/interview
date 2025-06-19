package main

import "sort"

func countDays(days int, meetings [][]int) int {
	var ret int
	d := 1
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][0] == meetings[j][0] {
			return meetings[i][1] < meetings[j][1]
		}
		return meetings[i][0] < meetings[j][0]
	})
	for _, m := range meetings {
		ret += max(m[0]-d, 0)
		d = max(d, m[1])
	}
	ret += days - d
	return ret
}
