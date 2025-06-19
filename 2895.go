package main

import "sort"

func minProcessingTime(processorTime []int, tasks []int) int {
	sort.Ints(processorTime)
	sort.Ints(tasks)
	var ret int
	for i, j := 0, len(tasks)-1; i < len(processorTime); i++ {
		ret = max(ret, processorTime[i]+tasks[j])
		j -= 4
	}
	return ret
}
