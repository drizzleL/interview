package main

func maxFreeTime2(eventTime int, startTime []int, endTime []int) int {
	beforeMax := make([]int, len(startTime))
	for i, before := 1, startTime[0]; i < len(startTime); i++ {
		beforeMax[i] = max(beforeMax[i-1], before)
		before = startTime[i] - endTime[i-1]
	}
	afterMax := make([]int, len(startTime))
	for i, after := len(startTime)-2, eventTime-endTime[len(endTime)-1]; i >= 0; i-- {
		afterMax[i] = max(afterMax[i+1], after)
		after = startTime[i+1] - endTime[i]
	}
	before, after := startTime[0], startTime[1]-endTime[0]
	var ret int
	for i := range startTime {
		start, end := startTime[i], endTime[i]
		gap := end - start
		if beforeMax[i] >= gap || afterMax[i] >= gap {
			ret = max(ret, after+before+gap)
		} else {
			ret = max(ret, after+before)
		}
		before = after
		if i < len(startTime)-2 {
			after = startTime[i+2] - endTime[i+1]
		} else {
			after = eventTime - endTime[len(endTime)-1]
		}
	}
	return ret
}
