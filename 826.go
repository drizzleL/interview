package main

import "sort"

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	var jobs [][2]int
	for i := range difficulty {
		jobs = append(jobs, [2]int{difficulty[i], profit[i]})
	}
	sort.Slice(jobs, func(i, j int) bool {
		if jobs[i][0] == jobs[j][0] {
			return jobs[i][1] < jobs[j][1]
		}
		return jobs[i][0] < jobs[j][0]
	})
	sort.Ints(worker)
	var ret int
	var maxProfit int
	var j int
	for _, ability := range worker {
		for j < len(jobs) && jobs[j][0] <= ability {
			maxProfit = max(maxProfit, jobs[j][1])
			j += 1
		}
		ret += maxProfit
	}
	return ret
}
