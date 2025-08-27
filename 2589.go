package main

import "sort"

func findMinimumTime2(tasks [][]int) int {
	var lines [2001]bool
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1] < tasks[j][1]
	})
	for i := 0; i < len(tasks); i++ {
		start, end, gap := tasks[i][0], tasks[i][1], tasks[i][2]
		for j := start; j <= end; j++ {
			if lines[j] {
				gap -= 1
				continue
			}
		}
		for i := end; gap > 0; i-- {
			if lines[i] {
				continue
			}
			lines[i] = true
			gap -= 1
		}
	}
	var ret int
	for _, l := range lines {
		if l {
			ret += 1
		}
	}
	return ret
}
