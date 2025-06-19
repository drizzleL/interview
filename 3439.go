package main

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	gaps := make([]int, len(startTime)+1)
	gaps[0] = startTime[0]
	gaps[len(gaps)-1] = eventTime - endTime[len(endTime)-1]
	for i := 1; i < len(gaps)-1; i++ {
		gaps[i] = startTime[i] - endTime[i-1]
	}
	presum := make([]int, len(gaps)+1)
	for i := 1; i < len(presum); i++ {
		presum[i] = presum[i-1] + gaps[i-1]
	}
	var ret int
	for i := 0; i < len(presum); i++ {
		end := min(i+k, len(gaps)-1)
		tmp := presum[end+1] - presum[i]
		ret = max(ret, tmp)
	}
	return ret
}
